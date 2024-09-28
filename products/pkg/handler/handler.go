package handler

import (
	"context"
	"strconv"

	"github.com/gofrs/uuid"
	pb "github.com/jijengeai/jijengeai/systems/products/pb/gen"
	"github.com/jijengeai/jijengeai/systems/products/pkg/db/models"
	svc "github.com/jijengeai/jijengeai/systems/products/pkg/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	pb.UnimplementedProductServiceServer
	service *svc.Service
	logger  *logrus.Logger
}

func NewHandler(svc *svc.Service, logger *logrus.Logger) *Handler {
	return &Handler{service: svc, logger: logger}
}

func (h *Handler) CreateFinancialInstitution(ctx context.Context, req *pb.CreateFinancialInstitutionRequest) (*pb.CreateFinancialInstitutionResponse, error) {
	h.logger.Infof("Creating financial institution: %s", req.Name)

	institution, err := h.service.CreateFinancialInstitution(req.Name)
	if err != nil {
		h.logger.Errorf("Failed to create financial institution: %v", err)
		return nil, status.Error(codes.Internal, "could not create financial institution")
	}

	return &pb.CreateFinancialInstitutionResponse{Institution: strconv.FormatUint(uint64(institution.ID), 10)}, nil
}

func (h *Handler) CreateCriteria(ctx context.Context, req *pb.CreateCriteriaRequest) (*pb.CreateCriteriaResponse, error) {
    h.logger.Infof("Creating criteria for institution ID: %s", req.Criteria.InstitutionId)

    institutionID, err := uuid.FromString(req.Criteria.InstitutionId)
    if err != nil {
        h.logger.Errorf("Invalid institution ID: %v", err)
        return nil, status.Error(codes.InvalidArgument, "invalid institution ID")
    }

    criteria := models.Criteria{
        NumberOfTransactions: int(req.Criteria.NumberOfTransactions),
        TotalAmount:         req.Criteria.TotalAmount,
        TaxPaid:             req.Criteria.TaxPaid,
        FrequentAmounts:     req.Criteria.FrequentAmounts,
        MinAmount:           req.Criteria.MinAmount,
        MaxAmount:           req.Criteria.MaxAmount,
    }

    createdCriteria, err := h.service.CreateCriteria(institutionID, criteria)
    if err != nil {
        h.logger.Errorf("Failed to create criteria: %v", err)
        return nil, status.Error(codes.Internal, "could not create criteria")
    }

    return &pb.CreateCriteriaResponse{CriteriaId: strconv.FormatUint(uint64(createdCriteria.ID), 10)}, nil
}


func (h *Handler) ListCriteriaByInstitutionId(ctx context.Context, req *pb.ListCriteriaByInstitutionIdRequest) (*pb.ListCriteriaByInstitutionIdResponse, error) {
	h.logger.Infof("Listing criteria for institution ID: %s", req.InstitutionId)

	institutionID, err := uuid.FromString(req.InstitutionId)
	if err != nil {
		h.logger.Errorf("Invalid institution ID: %v", err)
		return nil, status.Error(codes.InvalidArgument, "invalid institution ID")
	}

	criterias, err := h.service.ListCriteriaByInstitutionId(institutionID)
	if err != nil {
		h.logger.Errorf("Failed to list criteria: %v", err)
		return nil, status.Error(codes.Internal, "could not list criteria")
	}

	var responseCriterias []*pb.Criteria
	for _, c := range criterias {
		responseCriterias = append(responseCriterias, &pb.Criteria{
			InstitutionId:      c.InstitutionID.String(),
			NumberOfTransactions: int32(c.NumberOfTransactions),
			TotalAmount:        c.TotalAmount,
			TaxPaid:            c.TaxPaid,
			FrequentAmounts:    c.FrequentAmounts,
			MinAmount:          c.MinAmount,
			MaxAmount:          c.MaxAmount,
		})
	}

	return &pb.ListCriteriaByInstitutionIdResponse{Criterias: responseCriterias}, nil
}