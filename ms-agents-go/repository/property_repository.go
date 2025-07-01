package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"ms-agent/dto"
	"net/http"
)

type PropertyRepository interface {
	GetAllProperties() ([]dto.PropertyResponse, error)
	GetPropertyByID(id string) (*dto.PropertyResponse, error)
	GetMyProperties(accessToken string) ([]dto.PropertyResponse, error)
	CreateProperty(accessToken string, request dto.PropertyRequest) (dto.PropertyRequest, error)
}

type propertyRepository struct {
	propertiesServiceURL string
}

func NewPropertyRepository(propertiesServiceURL string) PropertyRepository {
	return &propertyRepository{propertiesServiceURL: propertiesServiceURL}
}

// obtener lista propiedades
func (pr *propertyRepository) GetAllProperties() ([]dto.PropertyResponse, error) {
	url := fmt.Sprintf("%s/api/v1/properties", pr.propertiesServiceURL)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to properties service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("properties service returned status: %d", resp.StatusCode)
	}

	var properties []dto.PropertyResponse
	if err := json.NewDecoder(resp.Body).Decode(&properties); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return properties, nil
}

// obtener propiedad por id
func (pr *propertyRepository) GetPropertyByID(id string) (*dto.PropertyResponse, error) {
	url := fmt.Sprintf("%s/api/v1/properties/%s", pr.propertiesServiceURL, id)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to properties service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("properties service returned status: %d", resp.StatusCode)
	}

	var properties *dto.PropertyResponse
	if err := json.NewDecoder(resp.Body).Decode(&properties); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return properties, nil
}

// obtener propiedades asignadas
func (pr *propertyRepository) GetMyProperties(accessToken string) ([]dto.PropertyResponse, error) {
	url := fmt.Sprintf("%s/api/v1/properties/me", pr.propertiesServiceURL)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to properties service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("properties service returned status: %d", resp.StatusCode)
	}

	var properties []dto.PropertyResponse
	if err := json.NewDecoder(resp.Body).Decode(&properties); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return properties, nil
}

// crear nueva propiedad
func (pr *propertyRepository) CreateProperty(accessToken string, request dto.PropertyRequest) (dto.PropertyRequest, error) {
	url := fmt.Sprintf("%s/api/v1/properties", pr.propertiesServiceURL)

	// Convertimos el dto a JSON
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return dto.PropertyRequest{}, fmt.Errorf("failed to marshal property request: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return dto.PropertyRequest{}, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dto.PropertyRequest{}, fmt.Errorf("failed to connect to properties service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return dto.PropertyRequest{}, fmt.Errorf("properties service returned status: %d", resp.StatusCode)
	}

	var createdProperty dto.PropertyRequest
	if err := json.NewDecoder(resp.Body).Decode(&createdProperty); err != nil {
		return dto.PropertyRequest{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return createdProperty, nil
}
