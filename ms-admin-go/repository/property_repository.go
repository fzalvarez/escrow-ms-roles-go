package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"ms-admin/dto"
	"net/http"
)

type PropertyRepository interface {
	GetPropertyByID(id string) (*dto.PropertyResponse, error)
	GetAllProperties() ([]dto.PropertyResponse, error)
	UpdateOwner(accessToken string, propertyID string, newOwnerID string) error
	AssignAgent(accessToken string, propertyID string, agentID string) error
	UpdateStatus(accessToken string, propertyID string, status string) error
}

type propertyRepository struct {
	propertiesServiceURL string
}

func NewPropertyRepository(propertiesServiceURL string) PropertyRepository {
	return &propertyRepository{propertiesServiceURL: propertiesServiceURL}
}

// lista de propiedades
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

// propiedad por id
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

// actualizar propietario a una propiedad
func (pr *propertyRepository) UpdateOwner(accessToken string, propertyID string, newOwnerID string) error {
	url := fmt.Sprintf("%s/api/v1/properties/%s/owner", pr.propertiesServiceURL, propertyID)

	payload := map[string]string{
		"ownerId": newOwnerID,
	}
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to create Patch request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to connect to properties service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("properties service returned unexpected status: %d", resp.StatusCode)
	}

	return nil
}

// asignar agente a una propiedad
func (pr *propertyRepository) AssignAgent(accessToken string, propertyID string, agentID string) error {
	url := fmt.Sprintf("%s/api/v1/properties/%s/agents", pr.propertiesServiceURL, propertyID)

	payload := map[string]string{
		"agent_id": agentID,
	}
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to create Patch request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to connect to properties service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("properties service returned unexpected status: %d", resp.StatusCode)
	}

	return nil
}

// actualizar status de una propiedad
func (pr *propertyRepository) UpdateStatus(accessToken string, propertyID string, status string) error {
	url := fmt.Sprintf("%s/api/v1/properties/%s/status", pr.propertiesServiceURL, propertyID)

	payload := map[string]string{
		"status": status,
	}
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to create Patch request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to connect to properties service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("properties service returned unexpected status: %d", resp.StatusCode)
	}

	return nil

}
