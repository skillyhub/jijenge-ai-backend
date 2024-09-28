package repository

import (
	"github.com/gofrs/uuid"
	"github.com/jijengeai/jijengeai/systems/products/pkg/db/models"
	"gorm.io/gorm"
)

type Repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
    return &Repository{db: db}
}



func (r *Repository) CreateFinancialInstitution(name string) (*models.FinancialInstitution, error) {
    institution := &models.FinancialInstitution{Name: name}
    if err := r.db.Create(institution).Error; err != nil {
        return nil, err
    }
    return institution, nil
}

func (r *Repository) CreateCriteria(institutionID uuid.UUID, criteria models.Criteria) (*models.Criteria, error) {
    criteria.InstitutionID = institutionID
    if err := r.db.Create(&criteria).Error; err != nil {
        return nil, err
    }
    return &criteria, nil
}

func (r *Repository) GetCriteriaById(id uuid.UUID) (*models.Criteria, error) {
    var criteria models.Criteria
    if err := r.db.Preload("FinancialInstitution").First(&criteria, id).Error; err != nil {
        return nil, err
    }
    return &criteria, nil
}
// GetCriteriaById retrieves a criteria by its UUID.

func (r *Repository) ListCriteriaByInstitutionId(institutionID uuid.UUID) ([]models.Criteria, error) {
    var criterias []models.Criteria
    if err := r.db.Where("institution_id = ?", institutionID).Find(&criterias).Error; err != nil {
        return nil, err
    }
    return criterias, nil
}