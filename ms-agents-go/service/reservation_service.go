package service

import (
	"ms-agent/dto"
	"ms-agent/repository"
)

type ReservationService interface {
	GetMyReservations(accessToken string) ([]dto.ReservationResponse, error)
	CreateRerservation(accessToken string, reservation *dto.ReservationRequest) (dto.ReservationResponse, error)
	UpdateStatus(accessToken string, reservationID string, status string) error
	GetReservationById(accessToken string, reservationID string) (*dto.ReservationResponse, error)
}

type reservationService struct {
	repo repository.ReservationRepository
}

func NewReservationService(repo repository.ReservationRepository) ReservationService {
	return &reservationService{repo: repo}
}

// obtener reservas por angete logueado
func (s *reservationService) GetMyReservations(accessToken string) ([]dto.ReservationResponse, error) {
	return s.repo.GetMyReservations(accessToken)
}

// crear reserva
func (s *reservationService) CreateRerservation(accessToken string, reservation *dto.ReservationRequest) (dto.ReservationResponse, error) {
	return s.repo.CreateRerservation(accessToken, reservation)
}

// actualizar status reserva
func (s *reservationService) UpdateStatus(accessToken, reservationID string, status string) error {
	statusReq := &dto.StatusRequest{Status: status}
	return s.repo.UpdateStatus(accessToken, reservationID, statusReq)
}

// obtener reserva por id
func (s *reservationService) GetReservationById(accessToken string, reservationID string) (*dto.ReservationResponse, error) {
	return s.repo.GetReservationById(accessToken, reservationID)
}
