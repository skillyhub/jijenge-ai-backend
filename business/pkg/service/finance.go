package service

import (
	"errors"

	"github.com/gofrs/uuid"
	"github.com/jijengeai/jijengeai/systems/business/pkg/db/models"
	repository "github.com/jijengeai/jijengeai/systems/business/pkg/db/repo"
)

var (
    ErrBusinessNotFound = errors.New("business not found")
)

type Service struct {
    repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
    return &Service{repo: repo}
}

func (s *Service) CreateBusiness(name, regID, email, phone string, businessType models.BusinessType) (*models.Business, error) {
    business := &models.Business{
        Name:  name,
        RegID: regID,
        Email: email,
        Phone: phone,
        Type:  businessType,
    }
    err := s.repo.CreateBusiness(business)
    if err != nil {
        return nil, err
    }
    return business, nil
}

func (s *Service) GetBusinessByID(id uuid.UUID) (*models.Business, error) {
    business, err := s.repo.GetBusinessByID(id)
    if err != nil {
        return nil, err
    }
    return business, nil
}

func (s *Service) UpdateBusiness(id uuid.UUID, name, email, phone string, businessType models.BusinessType) (*models.Business, error) {
    business, err := s.repo.GetBusinessByID(id)
    if err != nil {
       
        return nil, err
    }

    business.Name = name
    business.Email = email
    business.Phone = phone
    business.Type = businessType

    err = s.repo.UpdateBusiness(business)
    if err != nil {
        return nil, err
    }
    return business, nil
}

func (s *Service) DeleteBusiness(id uuid.UUID) error {
    err := s.repo.DeleteBusiness(id)
    if err != nil {
     
        return err
    }
    return nil
}


func (s *Service) SearchBusinesses(query string) ([]*models.Business, error) {
    return s.repo.SearchBusinesses(query)
}