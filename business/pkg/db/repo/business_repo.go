package repository

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/jijengeai/jijengeai/systems/business/pkg/db/models"
	"gorm.io/gorm"
)

type Repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
    return &Repository{db: db}
}

func (r *Repository) CreateBusiness(business *models.Business) error {
    return r.db.Create(business).Error
}

func (r *Repository) GetBusinessByID(id uuid.UUID) (*models.Business, error) {
    var business models.Business
    err := r.db.First(&business, "id = ?", id).Error
    if err != nil {
        return nil, err
    }
    return &business, nil
}

func (r *Repository) GetBusinessByRegID(regID string) (*models.Business, error) {
    var business models.Business
    err := r.db.First(&business, "reg_id = ?", regID).Error
    if err != nil {
        return nil, err
    }
    return &business, nil
}

func (r *Repository) UpdateBusiness(business *models.Business) error {
    return r.db.Save(business).Error
}

func (r *Repository) DeleteBusiness(id uuid.UUID) error {
    return r.db.Delete(&models.Business{}, "id = ?", id).Error
}

func (r *Repository) ListBusinesses(offset, limit int) ([]*models.Business, error) {
    var businesses []*models.Business
    err := r.db.Offset(offset).Limit(limit).Find(&businesses).Error
    return businesses, err
}

func (r *Repository) GetBusinessesByType(businessType models.BusinessType) ([]*models.Business, error) {
    var businesses []*models.Business
    err := r.db.Where("type = ?", businessType).Find(&businesses).Error
    return businesses, err
}

func (r *Repository) SearchBusinesses(query string) ([]*models.Business, error) {
    var businesses []*models.Business
    err := r.db.Where("name LIKE ? OR reg_id LIKE ?", "%"+query+"%", "%"+query+"%").Find(&businesses).Error
    return businesses, err
}

func (r *Repository) GetBusinessesCreatedBetween(start, end time.Time) ([]*models.Business, error) {
    var businesses []*models.Business
    err := r.db.Where("created_at BETWEEN ? AND ?", start, end).Find(&businesses).Error
    return businesses, err
}

func (r *Repository) CountBusinessesByType() (map[models.BusinessType]int64, error) {
    var results []struct {
        Type  models.BusinessType
        Count int64
    }
    err := r.db.Model(&models.Business{}).Select("type, count(*) as count").Group("type").Scan(&results).Error
    if err != nil {
        return nil, err
    }

    countMap := make(map[models.BusinessType]int64)
    for _, result := range results {
        countMap[result.Type] = result.Count
    }
    return countMap, nil
}

func (r *Repository) UpdateBusinessType(id uuid.UUID, newType models.BusinessType) error {
    return r.db.Model(&models.Business{}).Where("id = ?", id).Update("type", newType).Error
}

func (r *Repository) GetBusinessByEmail(email string) (*models.Business, error) {
    var business models.Business
    err := r.db.First(&business, "email = ?", email).Error
    if err != nil {
        return nil, err
    }
    return &business, nil
}