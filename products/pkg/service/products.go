package service

import (
	"github.com/gofrs/uuid"
	"github.com/jijengeai/jijengeai/systems/products/pkg/db/models"
	repository "github.com/jijengeai/jijengeai/systems/products/pkg/db/repo"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

// CreateFinancialInstitution creates a new financial institution in the database.
func (s *Service) CreateFinancialInstitution(name string) (*models.FinancialInstitution, error) {
	return s.repo.CreateFinancialInstitution(name)
}

func (s *Service) CreateCriteria(institutionID uuid.UUID, criteria models.Criteria) (*models.Criteria, error) {
	return s.repo.CreateCriteria(institutionID, criteria)
}

func (s *Service) ListCriteriaByInstitutionId(institutionID uuid.UUID) ([]models.Criteria, error) {
	return s.repo.ListCriteriaByInstitutionId(institutionID)
}