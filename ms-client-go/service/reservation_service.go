package service

import (
	"ms-clients/dto"
	"ms-clients/repository"
)

type ReservationService interface {
	GetMyReservations(accessToken string) ([]dto.ReservationResponse, error)
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

// obtener reserva por id
func (s *reservationService) GetReservationById(accessToken string, reservationID string) (*dto.ReservationResponse, error) {
	return s.repo.GetReservationById(accessToken, reservationID)
}
