package service

import (
	"ms-owners/dto"
	"ms-owners/repository"
)

type PropertyService interface {
	CreateProperty(accessToken string, request dto.PropertyRequest) (dto.PropertyRequest, error)
	GetPropertyByID(id string) (*dto.PropertyResponse, error)
	GetMyProperties(accessToken string) ([]dto.PropertyResponse, error)
	UpdatePropertyByID(accessToken string, id string, request dto.PropertyRequest) error
	DeletePropertyByID(accessToken string, id string) error
}

type propertyService struct {
	repo repository.PropertyRepository
}

func NewPropertyService(repo repository.PropertyRepository) PropertyService {
	return &propertyService{repo: repo}
}

// obtener propiedad por id
func (s *propertyService) GetPropertyByID(id string) (*dto.PropertyResponse, error) {
	return s.repo.GetPropertyByID(id)
}

// obtener mis propiedades
func (s *propertyService) GetMyProperties(accessToken string) ([]dto.PropertyResponse, error) {
	return s.repo.GetMyProperties(accessToken)
}

// crear  propiedad
func (s *propertyService) CreateProperty(accessToken string, request dto.PropertyRequest) (dto.PropertyRequest, error) {
	return s.repo.CreateProperty(accessToken, request)
}

// actualizar propiedad
func (s *propertyService) UpdatePropertyByID(accessToken string, id string, request dto.PropertyRequest) error {
	return s.repo.UpdatePropertyByID(accessToken, id, request)
}

// eliminar propiedad
func (s *propertyService) DeletePropertyByID(accessToken string, id string) error {
	return s.repo.DeletePropertyByID(accessToken, id)
}
