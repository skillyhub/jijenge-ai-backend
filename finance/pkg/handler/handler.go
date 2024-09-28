package handler

import (
	"context"
	"errors"

	"github.com/gofrs/uuid"
	pb "github.com/jijengeai/jijengeai/systems/finance/pb/gen"
	"github.com/jijengeai/jijengeai/systems/finance/pkg/db/models"
	svc "github.com/jijengeai/jijengeai/systems/finance/pkg/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Handler struct {
	pb.UnimplementedFinanceServiceServer
	service *svc.Service
	logger  *logrus.Logger
}

var (
	ErrInvalidAmount      = errors.New("amount must be greater than zero")
	ErrInvalidFinanceType = errors.New("invalid finance type")
	ErrInvalidBusinessID  = errors.New("invalid business ID")
)

func NewHandler(svc *svc.Service, logger *logrus.Logger) *Handler {
	return &Handler{service: svc, logger: logger}
}

func (h *Handler) CreateFinance(ctx context.Context, req *pb.CreateFinanceRequest) (*pb.CreateFinanceResponse, error) {
	h.logger.Infof("Creating finance for business ID: %s", req.BusinessId)

	businessID, err := uuid.FromString(req.BusinessId)
	if err != nil {
		h.logger.Error("Invalid business ID")
		return nil, ErrInvalidBusinessID
	}

	financeType := models.FinanceType(req.FinanceType)
	if err := financeType.IsValid(); err != nil {
		h.logger.Errorf("Invalid finance type: %v", err)
		return nil, ErrInvalidFinanceType
	}

	finance, err := h.service.CreateFinance(businessID, financeType, req.InitialAmount)
	if err != nil {
		h.logger.Errorf("Failed to create finance: %v", err)
		return nil, err
	}

	h.logger.Infof("Finance created successfully for business ID: %s", req.BusinessId)
	return &pb.CreateFinanceResponse{
		FinanceId: finance.Id.String(),
		Success:   true,
	}, nil
}

func (h *Handler) AddTransaction(ctx context.Context, req *pb.AddTransactionRequest) (*pb.AddTransactionResponse, error) {
	h.logger.Infof("Adding transaction for business ID: %s", req.BusinessId)

	businessID, err := uuid.FromString(req.BusinessId)
	if err != nil {
		h.logger.Error("Invalid business ID")
		return nil, ErrInvalidBusinessID
	}

	if req.Amount <= 0 {
		h.logger.Error("Invalid amount: must be greater than zero")
		return nil, ErrInvalidAmount
	}

	financeType := models.FinanceType(req.FinanceType)
	if err := financeType.IsValid(); err != nil {
		h.logger.Errorf("Invalid finance type: %v", err)
		return nil, ErrInvalidFinanceType
	}

	err = h.service.CreateTransaction(businessID, financeType, req.Amount, req.Description)
	if err != nil {
		h.logger.Errorf("Failed to add transaction: %v", err)
		return nil, err
	}

	h.logger.Infof("Transaction added successfully for business ID: %s", req.BusinessId)
	return &pb.AddTransactionResponse{Success: true}, nil
}

func (h *Handler) GetTransactions(ctx context.Context, req *pb.GetTransactionsRequest) (*pb.GetTransactionsResponse, error) {
	h.logger.Infof("Fetching transactions for business ID: %s", req.BusinessId)

	businessID, err := uuid.FromString(req.BusinessId)
	if err != nil {
		h.logger.Error("Invalid business ID")
		return nil, ErrInvalidBusinessID
	}

	var transactions []models.Transaction
	if req.FinanceType != "" {
		financeType := models.FinanceType(req.FinanceType)
		if err := financeType.IsValid(); err != nil {
			h.logger.Errorf("Invalid finance type: %v", err)
			return nil, ErrInvalidFinanceType
		}
		transactions, err = h.service.GetTransactionsByFinanceType(businessID, financeType)
	} else {
		transactions, err = h.service.GetTransactions(businessID)
	}

	if err != nil {
		h.logger.Errorf("Failed to get transactions: %v", err)
		return nil, err
	}

	pbTransactions := make([]*pb.Transaction, len(transactions))
	for i, t := range transactions {
		pbTransactions[i] = &pb.Transaction{
			Id:          t.Id.String(),
			FinanceType: string(t.FinanceType),
			Amount:      t.Amount,
			Description: t.Description,
			Date:        timestamppb.New(t.Date),
		}
	}

	h.logger.Infof("Successfully fetched transactions for business ID: %s", req.BusinessId)
	return &pb.GetTransactionsResponse{
		Transactions: pbTransactions,
	}, nil
}

func (h *Handler) GetFinancialSummary(ctx context.Context, req *pb.GetFinancialSummaryRequest) (*pb.GetFinancialSummaryResponse, error) {
	h.logger.Infof("Generating financial summary for business ID: %s", req.BusinessId)

	businessID, err := uuid.FromString(req.BusinessId)
	if err != nil {
		h.logger.Error("Invalid business ID")
		return nil, ErrInvalidBusinessID
	}

	summary, err := h.service.GenerateFinancialSummary(businessID)
	if err != nil {
		h.logger.Errorf("Failed to generate financial summary: %v", err)
		return nil, err
	}

	latestTransactions := make([]*pb.Transaction, len(summary["latestTransactions"].([]models.Transaction)))
	for i, t := range summary["latestTransactions"].([]models.Transaction) {
		latestTransactions[i] = &pb.Transaction{
			Id:          t.Id.String(),
			FinanceType: string(t.FinanceType),
			Amount:      t.Amount,
			Description: t.Description,
			Date:        timestamppb.New(t.Date),
		}
	}

	h.logger.Infof("Successfully generated financial summary for business ID: %s", req.BusinessId)
	return &pb.GetFinancialSummaryResponse{
		TotalBalance:         summary["totalBalance"].(float64),
		MomoBalance:          summary["momoBalance"].(float64),
		BankBalance:          summary["bankBalance"].(float64),
		MomoTransactionCount: summary["momoTransactionCount"].(int64),
		BankTransactionCount: summary["bankTransactionCount"].(int64),
		LatestTransactions:   latestTransactions,
	}, nil
}

func (h *Handler) TransferBetweenFinances(ctx context.Context, req *pb.TransferBetweenFinancesRequest) (*pb.TransferBetweenFinancesResponse, error) {
	h.logger.Infof("Transferring between finances for business ID: %s", req.BusinessId)

	businessID, err := uuid.FromString(req.BusinessId)
	if err != nil {
		h.logger.Error("Invalid business ID")
		return nil, ErrInvalidBusinessID
	}

	if req.Amount <= 0 {
		h.logger.Error("Invalid amount: must be greater than zero")
		return nil, ErrInvalidAmount
	}

	fromType := models.FinanceType(req.FromFinanceType)
	toType := models.FinanceType(req.ToFinanceType)
	if err := fromType.IsValid(); err != nil {
		h.logger.Error("Invalid finance type")
		return nil, ErrInvalidFinanceType
	}

	err = h.service.TransferBetweenFinances(businessID, fromType, toType, req.Amount)
	if err != nil {
		h.logger.Errorf("Failed to transfer between finances: %v", err)
		return nil, err
	}

	h.logger.Infof("Successfully transferred between finances for business ID: %s", req.BusinessId)
	return &pb.TransferBetweenFinancesResponse{Success: true}, nil
}