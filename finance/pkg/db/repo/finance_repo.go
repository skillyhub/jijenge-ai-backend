package repository

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/jijengeai/jijengeai/systems/finance/pkg/db/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}


func (r *Repository) GetFinanceByID(id uuid.UUID) (*models.Finance, error) {
	var finance models.Finance
	result := r.db.First(&finance, "id = ?", id)
	return &finance, result.Error
}

func (r *Repository) GetFinanceByBusinessID(businessID uuid.UUID, financeType models.FinanceType) (*models.Finance, error) {
	var finance models.Finance
	result := r.db.Where("business_id = ? AND type = ?", businessID, financeType).First(&finance)
	return &finance, result.Error
}

func (r *Repository) CreateFinance(finance *models.Finance) error {
	return r.db.Create(finance).Error
}

func (r *Repository) UpdateFinance(finance *models.Finance) error {
	return r.db.Save(finance).Error
}

func (r *Repository) DeleteFinance(id uuid.UUID) error {
	return r.db.Delete(&models.Finance{}, "id = ?", id).Error
}

func (r *Repository) ListFinancesByBusinessID(businessID uuid.UUID) ([]models.Finance, error) {
	var finances []models.Finance
	result := r.db.Where("business_id = ?", businessID).Find(&finances)
	return finances, result.Error
}


func (r *Repository) GetTransactionByID(id uuid.UUID) (*models.Transaction, error) {
	var transaction models.Transaction
	result := r.db.First(&transaction, "id = ?", id)
	return &transaction, result.Error
}

func (r *Repository) CreateTransaction(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *Repository) UpdateTransaction(transaction *models.Transaction) error {
	return r.db.Save(transaction).Error
}

func (r *Repository) DeleteTransaction(id uuid.UUID) error {
	return r.db.Delete(&models.Transaction{}, "id = ?", id).Error
}

func (r *Repository) ListTransactionsByBusinessID(businessID uuid.UUID) ([]models.Transaction, error) {
	var transactions []models.Transaction
	result := r.db.Where("business_id = ?", businessID).Find(&transactions)
	return transactions, result.Error
}

func (r *Repository) ListTransactionsByFinanceType(businessID uuid.UUID, financeType models.FinanceType) ([]models.Transaction, error) {
	var transactions []models.Transaction
	result := r.db.Where("business_id = ? AND finance_type = ?", businessID, financeType).Find(&transactions)
	return transactions, result.Error
}

func (r *Repository) ListTransactionsByDateRange(businessID uuid.UUID, startDate, endDate time.Time) ([]models.Transaction, error) {
	var transactions []models.Transaction
	result := r.db.Where("business_id = ? AND date BETWEEN ? AND ?", businessID, startDate, endDate).Find(&transactions)
	return transactions, result.Error
}


func (r *Repository) GetTotalBalance(businessID uuid.UUID) (float64, error) {
	var totalBalance float64
	result := r.db.Model(&models.Finance{}).Where("business_id = ?", businessID).Select("SUM(amount)").Scan(&totalBalance)
	return totalBalance, result.Error
}

func (r *Repository) GetBalanceByFinanceType(businessID uuid.UUID, financeType models.FinanceType) (float64, error) {
	var balance float64
	result := r.db.Model(&models.Finance{}).Where("business_id = ? AND type = ?", businessID, financeType).Select("amount").Scan(&balance)
	return balance, result.Error
}

func (r *Repository) TransferBetweenFinances(fromFinanceID, toFinanceID uuid.UUID, amount float64) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Finance{}).Where("id = ?", fromFinanceID).UpdateColumn("amount", gorm.Expr("amount - ?", amount)).Error; err != nil {
			return err
		}

		if err := tx.Model(&models.Finance{}).Where("id = ?", toFinanceID).UpdateColumn("amount", gorm.Expr("amount + ?", amount)).Error; err != nil {
			return err
		}

		return nil
	})
}


func (r *Repository) GetTotalTransactionAmount(businessID uuid.UUID, financeType models.FinanceType) (float64, error) {
	var totalAmount float64
	result := r.db.Model(&models.Transaction{}).Where("business_id = ? AND finance_type = ?", businessID, financeType).Select("SUM(amount)").Scan(&totalAmount)
	return totalAmount, result.Error
}

func (r *Repository) GetTransactionCountByFinanceType(businessID uuid.UUID, financeType models.FinanceType) (int64, error) {
	var count int64
	result := r.db.Model(&models.Transaction{}).Where("business_id = ? AND finance_type = ?", businessID, financeType).Count(&count)
	return count, result.Error
}

func (r *Repository) GetLatestTransactions(businessID uuid.UUID, limit int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	result := r.db.Where("business_id = ?", businessID).Order("date DESC").Limit(limit).Find(&transactions)
	return transactions, result.Error
}