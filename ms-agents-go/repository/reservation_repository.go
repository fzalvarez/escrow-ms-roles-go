package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"ms-agent/dto"
	"net/http"
)

type ReservationRepository interface {
	GetMyReservations(accessToken string) ([]dto.ReservationResponse, error)
	CreateRerservation(accessToken string, reservation *dto.ReservationRequest) (dto.ReservationResponse, error)
	UpdateStatus(accessToken string, reservationID string, status *dto.StatusRequest) error
	GetReservationById(accessToken string, reservationID string) (*dto.ReservationResponse, error)
}

type reservationRepository struct {
	reservationServiceURL string
}

func NewReservationRepository(reservationServiceURL string) ReservationRepository {
	return &reservationRepository{reservationServiceURL: reservationServiceURL}
}

// obtener reservas por angete logueado
func (r *reservationRepository) GetMyReservations(accessToken string) ([]dto.ReservationResponse, error) {
	url := fmt.Sprintf("%s/api/v1/reservation/agent/me", r.reservationServiceURL)

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

// crear reserva
func (r *reservationRepository) CreateRerservation(accessToken string, reservation *dto.ReservationRequest) (dto.ReservationResponse, error) {
	url := fmt.Sprintf("%s/api/v1/reservation/", r.reservationServiceURL)

	jsonBody, err := json.Marshal(reservation)
	if err != nil {
		return dto.ReservationResponse{}, fmt.Errorf("failed to marshal property request: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return dto.ReservationResponse{}, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dto.ReservationResponse{}, fmt.Errorf("failed to connect to properties service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return dto.ReservationResponse{}, fmt.Errorf("properties service returned status: %d", resp.StatusCode)
	}

	var createdReservation dto.ReservationResponse
	if err := json.NewDecoder(resp.Body).Decode(&createdReservation); err != nil {
		return dto.ReservationResponse{}, fmt.Errorf("failed to decode response: %w", err)
	}
	return createdReservation, nil
}

// actualizar status reserva
func (r *reservationRepository) UpdateStatus(accessToken string, reservationID string, status *dto.StatusRequest) error {
	url := fmt.Sprintf("%s/api/v1/reservation/%s", r.reservationServiceURL, reservationID)

	jsonBody, err := json.Marshal(status)
	if err != nil {
		return fmt.Errorf("failed to marshal reservation request: %w", err)
	}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to connect to properties service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("reservation service returned status: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil

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
