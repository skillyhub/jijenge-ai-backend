package service

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jijengeai/jijengeai/systems/finance/pkg/db/models"
	repository "github.com/jijengeai/jijengeai/systems/finance/pkg/db/repo"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

// Finance-related methods

func (s *Service) CreateFinance(businessID uuid.UUID, financeType models.FinanceType, initialAmount float64) (*models.Finance, error) {
	financeID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	finance := &models.Finance{
		Id:         financeID,
		BusinessID: businessID,
		Type:       financeType,
		Amount:     initialAmount,
	}

	err = s.repo.CreateFinance(finance)
	if err != nil {
		return nil, err
	}

	return finance, nil
}

func (s *Service) GetFinance(businessID uuid.UUID, financeType models.FinanceType) (*models.Finance, error) {
	return s.repo.GetFinanceByBusinessID(businessID, financeType)
}

func (s *Service) UpdateFinanceAmount(businessID uuid.UUID, financeType models.FinanceType, amount float64) error {
	finance, err := s.repo.GetFinanceByBusinessID(businessID, financeType)
	if err != nil {
		return err
	}

	finance.Amount = amount
	return s.repo.UpdateFinance(finance)
}

func (s *Service) AddMoneyToFinance(businessID uuid.UUID, financeType models.FinanceType, amount float64) error {
	finance, err := s.repo.GetFinanceByBusinessID(businessID, financeType)
	if err != nil {
		return err
	}

	finance.Amount += amount
	return s.repo.UpdateFinance(finance)
}

func (s *Service) TransferBetweenFinances(businessID uuid.UUID, fromType, toType models.FinanceType, amount float64) error {
	fromFinance, err := s.repo.GetFinanceByBusinessID(businessID, fromType)
	if err != nil {
		return err
	}

	toFinance, err := s.repo.GetFinanceByBusinessID(businessID, toType)
	if err != nil {
		return err
	}

	if fromFinance.Amount < amount {
		return errors.New("insufficient funds")
	}

	return s.repo.TransferBetweenFinances(fromFinance.Id, toFinance.Id, amount)
}

func (s *Service) GetTotalBalance(businessID uuid.UUID) (float64, error) {
	return s.repo.GetTotalBalance(businessID)
}

// Transaction-related methods

func (s *Service) CreateTransaction(businessID uuid.UUID, financeType models.FinanceType, amount float64, description string) error {
	transactionID, err := uuid.NewV4()
	if err != nil {
		return err
	}

	transaction := &models.Transaction{
		Id:          transactionID,
		BusinessID:  businessID,
		FinanceType: financeType,
		Amount:      amount,
		Description: description,
		Date:        time.Now(),
	}

	err = s.repo.CreateTransaction(transaction)
	if err != nil {
		return err
	}

	// Update the corresponding finance
	return s.AddMoneyToFinance(businessID, financeType, amount)
}

func (s *Service) GetTransactions(businessID uuid.UUID) ([]models.Transaction, error) {
	return s.repo.ListTransactionsByBusinessID(businessID)
}

func (s *Service) GetTransactionsByFinanceType(businessID uuid.UUID, financeType models.FinanceType) ([]models.Transaction, error) {
	return s.repo.ListTransactionsByFinanceType(businessID, financeType)
}

func (s *Service) GetTransactionsByDateRange(businessID uuid.UUID, startDate, endDate time.Time) ([]models.Transaction, error) {
	return s.repo.ListTransactionsByDateRange(businessID, startDate, endDate)
}

func (s *Service) GetTotalTransactionAmount(businessID uuid.UUID, financeType models.FinanceType) (float64, error) {
	return s.repo.GetTotalTransactionAmount(businessID, financeType)
}

func (s *Service) GetTransactionCountByFinanceType(businessID uuid.UUID, financeType models.FinanceType) (int64, error) {
	return s.repo.GetTransactionCountByFinanceType(businessID, financeType)
}

func (s *Service) GetLatestTransactions(businessID uuid.UUID, limit int) ([]models.Transaction, error) {
	return s.repo.GetLatestTransactions(businessID, limit)
}


func (s *Service) GenerateFinancialSummary(businessID uuid.UUID) (map[string]interface{}, error) {
	totalBalance, err := s.GetTotalBalance(businessID)
	if err != nil {
		return nil, err
	}

	momoBalance, err := s.repo.GetBalanceByFinanceType(businessID, models.FinanceTypeMomo)
	if err != nil {
		return nil, err
	}

	bankBalance, err := s.repo.GetBalanceByFinanceType(businessID, models.FinanceTypeBank)
	if err != nil {
		return nil, err
	}

	momoTransactionCount, err := s.GetTransactionCountByFinanceType(businessID, models.FinanceTypeMomo)
	if err != nil {
		return nil, err
	}

	bankTransactionCount, err := s.GetTransactionCountByFinanceType(businessID, models.FinanceTypeBank)
	if err != nil {
		return nil, err
	}

	latestTransactions, err := s.GetLatestTransactions(businessID, 5)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"totalBalance":         totalBalance,
		"momoBalance":          momoBalance,
		"bankBalance":          bankBalance,
		"momoTransactionCount": momoTransactionCount,
		"bankTransactionCount": bankTransactionCount,
		"latestTransactions":   latestTransactions,
	}, nil
}