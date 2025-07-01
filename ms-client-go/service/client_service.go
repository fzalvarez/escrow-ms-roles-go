package service

import (
	"ms-clients/dto"
	"ms-clients/repository"
)

type UserService interface {
	GetMyProfile(accessToken string) (dto.User, error)
	UpdateProfile(accessToken string, profile dto.UpdateProfileRequest) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// obtener mi perfil
func (s *userService) GetMyProfile(accessToken string) (dto.User, error) {
	return s.repo.GetMyProfile(accessToken)
}

// actualizar mi perfil
func (s *userService) UpdateProfile(accessToken string, profile dto.UpdateProfileRequest) error {
	return s.repo.UpdateProfile(accessToken, profile)
}
