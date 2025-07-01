package repository

import (
	"encoding/json"
	"fmt"
	"ms-clients/dto"
	"net/http"
)

type ReservationRepository interface {
	GetMyReservations(accessToken string) ([]dto.ReservationResponse, error)
	GetReservationById(accessToken string, reservationID string) (*dto.ReservationResponse, error)
}

type reservationRepository struct {
	reservationServiceURL string
}

func NewReservationRepository(reservationServiceURL string) ReservationRepository {
	return &reservationRepository{reservationServiceURL: reservationServiceURL}
}

// obtener reservas por cliente logueado
func (r *reservationRepository) GetMyReservations(accessToken string) ([]dto.ReservationResponse, error) {
	url := fmt.Sprintf("%s/api/v1/reservation/client/me", r.reservationServiceURL)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []dto.ReservationResponse{}, fmt.Errorf("failed to create request:%w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []dto.ReservationResponse{}, fmt.Errorf("failed to connect to properties service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return []dto.ReservationResponse{}, fmt.Errorf("properties service returned unexpected status: %d", resp.StatusCode)
	}

	var reservations []dto.ReservationResponse
	err = json.NewDecoder(resp.Body).Decode(&reservations)
	if err != nil {
		return []dto.ReservationResponse{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return reservations, nil
}

// obtener reserva por id
func (r *reservationRepository) GetReservationById(accessToken string, reservationID string) (*dto.ReservationResponse, error) {
	url := fmt.Sprintf("%s/api/v1/reservation/%s", r.reservationServiceURL, reservationID)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to reservation service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("reservation service returned status: %d", resp.StatusCode)
	}

	var reservation dto.ReservationResponse
	if err := json.NewDecoder(resp.Body).Decode(&reservation); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &reservation, nil
}
