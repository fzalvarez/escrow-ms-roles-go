package service

import (
	"ms-admin/dto"
	"ms-admin/repository"
)

type UserService interface {
	GetClient(accessToken string) ([]dto.User, error)
	GetClientById(accessToken string, clientID string) (dto.User, error)
	GetOwner(accessToken string) ([]dto.User, error)
	GetOwnerById(accessToken string, clientID string) (dto.User, error)
	GetAgent(accessToken string) ([]dto.User, error)
	GetAgentById(accessToken string, clientID string) (dto.User, error)
	EnableAgent(accessToken string, agentID string) error
	DisableAgent(accessToken string, agenteID string) error
	EnableClient(accessToken string, clientID string) error
	DisableClient(accessToken string, clientID string) error
	EnableOwner(accessToken string, ownerID string) error
	DisableOwner(accessToken string, ownerID string) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// obtener lista clientes
func (s *userService) GetClient(accessToken string) ([]dto.User, error) {
	return s.repo.GetClient(accessToken)
}

// obtener cliente por id
func (s *userService) GetClientById(accessToken string, clientID string) (dto.User, error) {
	return s.repo.GetClientById(accessToken, clientID)
}

// obtener lista propietarios
func (s *userService) GetOwner(accessToken string) ([]dto.User, error) {
	return s.repo.GetOwner(accessToken)
}

// obtener propietario por id
func (s *userService) GetOwnerById(accessToken string, clientID string) (dto.User, error) {
	return s.repo.GetOwnerById(accessToken, clientID)
}

// obtener lista agentes
func (s *userService) GetAgent(accessToken string) ([]dto.User, error) {
	return s.repo.GetAgent(accessToken)
}

// obtener agente por id
func (s *userService) GetAgentById(accessToken string, clientID string) (dto.User, error) {
	return s.repo.GetAgentById(accessToken, clientID)
}

// deshabilitar agente por id
func (s *userService) DisableAgent(accessToken string, agenteID string) error {
	return s.repo.DisableAgent(accessToken, agenteID)
}

// deshabilitar cliente por id
func (s *userService) DisableClient(accessToken string, clientID string) error {
	return s.repo.DisableClient(accessToken, clientID)
}

// deshabilitar propietario por id
func (s *userService) DisableOwner(accessToken string, ownerID string) error {
	return s.repo.DisableOwner(accessToken, ownerID)
}

// habilitar agente por id
func (s *userService) EnableAgent(accessToken string, agentID string) error {
	return s.repo.EnableAgent(accessToken, agentID)
}

// habilitar cliente por id
func (s *userService) EnableClient(accessToken string, clientID string) error {
	return s.repo.EnableClient(accessToken, clientID)
}

// habilitar propietario por id
func (s *userService) EnableOwner(accessToken string, ownerID string) error {
	return s.repo.EnableOwner(accessToken, ownerID)
}
