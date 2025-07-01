package service

import (
	"ms-admin/dto"
	"ms-admin/repository"
)

type PropertyService interface {
	ListProperties() ([]dto.PropertyResponse, error)
	GetPropertyByID(id string) (*dto.PropertyResponse, error)
	UpdateOwner(accessToken string, propertyID string, newOwnerID string) error
	AssignAgent(accessToken string, propertyID string, agentID string) error
	UpdateStatus(accessToken string, propertyID string, status string) error
}

type propertyService struct {
	repo repository.PropertyRepository
}

func NewPropertyService(repo repository.PropertyRepository) PropertyService {
	return &propertyService{repo: repo}
}

// lista de propiedades
func (s *propertyService) ListProperties() ([]dto.PropertyResponse, error) {
	return s.repo.GetAllProperties()
}

// propiedad por id
func (s *propertyService) GetPropertyByID(id string) (*dto.PropertyResponse, error) {
	return s.repo.GetPropertyByID(id)
}

// actualizar propietario de una propiedad
func (s *propertyService) UpdateOwner(accessToken string, propertyID string, newOwnerID string) error {
	return s.repo.UpdateOwner(accessToken, propertyID, newOwnerID)
}

// asignar agente a una propiedad
func (s *propertyService) AssignAgent(accessToken string, propertyID string, agentID string) error {
	return s.repo.AssignAgent(accessToken, propertyID, agentID)
}

// actualizar status de una propiedad
func (s *propertyService) UpdateStatus(accessToken string, propertyID string, status string) error {
	return s.repo.UpdateStatus(accessToken, propertyID, status)
}
