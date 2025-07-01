package service

import (
	"ms-agent/dto"
	"ms-agent/repository"
)

type PropertyService interface {
	ListProperties() ([]dto.PropertyResponse, error)
	GetPropertyByID(id string) (*dto.PropertyResponse, error)
	GetMyProperties(accessToken string) ([]dto.PropertyResponse, error)
	CreateProperty(accessToken string, request dto.PropertyRequest) (dto.PropertyRequest, error)
}

type propertyService struct {
	repo repository.PropertyRepository
}

func NewPropertyService(repo repository.PropertyRepository) PropertyService {
	return &propertyService{repo: repo}
}

// obtener lista propiedades
func (s *propertyService) ListProperties() ([]dto.PropertyResponse, error) {
	return s.repo.GetAllProperties()
}

// obtener propiedad por id
func (s *propertyService) GetPropertyByID(id string) (*dto.PropertyResponse, error) {
	return s.repo.GetPropertyByID(id)
}

// obtener propiedades asignadas
func (s *propertyService) GetMyProperties(accessToken string) ([]dto.PropertyResponse, error) {
	return s.repo.GetMyProperties(accessToken)
}

// crear nueva propiedad
func (s *propertyService) CreateProperty(accessToken string, request dto.PropertyRequest) (dto.PropertyRequest, error) {
	return s.repo.CreateProperty(accessToken, request)
}
