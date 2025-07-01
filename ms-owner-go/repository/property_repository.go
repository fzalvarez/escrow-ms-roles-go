package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"ms-owners/dto"
	"net/http"
)

type PropertyRepository interface {
	GetPropertyByID(id string) (*dto.PropertyResponse, error)
	GetMyProperties(accessToken string) ([]dto.PropertyResponse, error)
	CreateProperty(accessToken string, request dto.PropertyRequest) (dto.PropertyRequest, error)
	UpdatePropertyByID(accessToken string, id string, request dto.PropertyRequest) error
	DeletePropertyByID(accessToken string, id string) error
}

type propertyRepository struct {
	propertiesServiceURL string
}

func NewPropertyRepository(propertiesServiceURL string) PropertyRepository {
	return &propertyRepository{propertiesServiceURL: propertiesServiceURL}
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

// obtener mis propiedades
func (pr *propertyRepository) GetMyProperties(accessToken string) ([]dto.PropertyResponse, error) {
	url := fmt.Sprintf("%s/api/v1/properties/owner", pr.propertiesServiceURL)

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

// crear  propiedad
func (pr *propertyRepository) CreateProperty(accessToken string, request dto.PropertyRequest) (dto.PropertyRequest, error) {
	url := fmt.Sprintf("%s/api/v1/properties", pr.propertiesServiceURL)

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

// actualizar propiedad
func (pr *propertyRepository) UpdatePropertyByID(accessToken string, id string, request dto.PropertyRequest) error {
	url := fmt.Sprintf("%s/api/v1/properties/%s", pr.propertiesServiceURL, id)
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	req, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", accessToken)
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("update failed, status=%d", resp.StatusCode)
	}
	return nil

}

// eliminar propiedad
func (pr *propertyRepository) DeletePropertyByID(accessToken string, id string) error {
	url := fmt.Sprintf("%s/api/v1/properties/%s", pr.propertiesServiceURL, id)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create delete request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send delete request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("delete failed, status=%d", resp.StatusCode)
	}

	return nil
}
