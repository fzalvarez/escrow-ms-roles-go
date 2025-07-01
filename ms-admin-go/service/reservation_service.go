package service

import (
	"ms-admin/dto"
	"ms-admin/repository"
)

type ReservationService interface {
	CreateRerservation(accessToken string, reservation *dto.ReservationRequest) (dto.ReservationResponse, error)
	UpdateStatus(accessToken string, reservationID string, status string) error
	GetReservationById(accessToken string, reservationID string) (*dto.ReservationResponse, error)
	FindAll(accessToken string) ([]dto.ReservationResponse, error)
	GetReservationsByClientId(accessToken string, clientID string) ([]dto.ReservationResponse, error)
	GetReservationsByAgentId(accessToken string, agentID string) ([]dto.ReservationResponse, error)
}

type reservationService struct {
	repo repository.ReservationRepository
}

func NewReservationService(repo repository.ReservationRepository) ReservationService {
	return &reservationService{repo: repo}
}

// crear reserva
func (s *reservationService) CreateRerservation(accessToken string, reservation *dto.ReservationRequest) (dto.ReservationResponse, error) {
	return s.repo.CreateRerservation(accessToken, reservation)
}

// actualizar status reserva
func (s *reservationService) UpdateStatus(accessToken, reservationID string, status string) error {
	statusReq := &dto.Status_Request{Status: status}
	return s.repo.UpdateStatus(accessToken, reservationID, statusReq)
}

// obtener reserva por id
func (s *reservationService) GetReservationById(accessToken string, reservationID string) (*dto.ReservationResponse, error) {
	return s.repo.GetReservationById(accessToken, reservationID)
}

// obtener todas las reservas
func (s *reservationService) FindAll(accessToken string) ([]dto.ReservationResponse, error) {
	return s.repo.FindAll(accessToken)
}

// obtener reservas por client id
func (s *reservationService) GetReservationsByClientId(accessToken string, clientID string) ([]dto.ReservationResponse, error) {
	return s.repo.GetReservationsByClientId(accessToken, clientID)
}

// obtener reservas por agent id
func (s *reservationService) GetReservationsByAgentId(accessToken string, agentID string) ([]dto.ReservationResponse, error) {
	return s.repo.GetReservationsByAgentId(accessToken, agentID)
}
