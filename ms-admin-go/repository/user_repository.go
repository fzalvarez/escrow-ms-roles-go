package repository

import (
	"encoding/json"
	"fmt"
	"ms-admin/dto"
	"net/http"
)

type UserRepository interface {
	GetClient(accessToken string) ([]dto.User, error)
	GetClientById(accessToken string, clientID string) (dto.User, error)
	GetOwner(accessToken string) ([]dto.User, error)
	GetOwnerById(accessToken string, ownerID string) (dto.User, error)
	GetAgent(accessToken string) ([]dto.User, error)
	GetAgentById(accessToken string, ownerID string) (dto.User, error)
	EnableAgent(accessToken string, agentID string) error
	DisableAgent(accessToken string, agenteID string) error
	EnableClient(accessToken string, clientID string) error
	DisableClient(accessToken string, clientID string) error
	EnableOwner(accessToken string, ownerID string) error
	DisableOwner(accessToken string, ownerID string) error
}

type userRepository struct {
	userServiceURL string
}

func NewUserRepository(userServiceURL string) UserRepository {
	return &userRepository{userServiceURL: userServiceURL}
}

// obtener lista clientes
func (ur *userRepository) GetClient(accessToken string) ([]dto.User, error) {
	url := fmt.Sprintf("%s/users/by-role?role=CLIENT", ur.userServiceURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []dto.User{}, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []dto.User{}, fmt.Errorf("failed to connect to user service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
	}

	var clients []dto.User
	if err := json.NewDecoder(resp.Body).Decode(&clients); err != nil {
		return nil, fmt.Errorf("failed to decoded response: %w", err)
	}
	return clients, nil
}

// obtener cliente por id
func (ur *userRepository) GetClientById(accessToken string, clientID string) (dto.User, error) {
	url := fmt.Sprintf("%s/users/by-role/%s?role=CLIENT", ur.userServiceURL, clientID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return dto.User{}, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", accessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dto.User{}, fmt.Errorf("failed to connect to user service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
	}

	var user dto.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return dto.User{}, fmt.Errorf("failed to decoded response: %w", err)
	}
	return user, nil
}

// obtener lista propietarios
func (ur *userRepository) GetOwner(accessToken string) ([]dto.User, error) {
	url := fmt.Sprintf("%s/users/by-role?role=OWNER", ur.userServiceURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []dto.User{}, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []dto.User{}, fmt.Errorf("failed to connect to user service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
	}

	var owners []dto.User
	if err := json.NewDecoder(resp.Body).Decode(&owners); err != nil {
		return nil, fmt.Errorf("failed to decoded response: %w", err)
	}
	return owners, nil
}

// obtener propietario por id
func (ur *userRepository) GetOwnerById(accessToken string, ownerID string) (dto.User, error) {
	url := fmt.Sprintf("%s/users/by-role/%s?role=OWNER", ur.userServiceURL, ownerID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return dto.User{}, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", accessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dto.User{}, fmt.Errorf("failed to connect to user service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
	}

	var user dto.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return dto.User{}, fmt.Errorf("failed to decoded response: %w", err)
	}
	return user, nil
}

// obtener lista agentes
func (ur *userRepository) GetAgent(accessToken string) ([]dto.User, error) {
	url := fmt.Sprintf("%s/users/by-role?role=AGENT", ur.userServiceURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []dto.User{}, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []dto.User{}, fmt.Errorf("failed to connect to user service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
	}

	var agents []dto.User
	if err := json.NewDecoder(resp.Body).Decode(&agents); err != nil {
		return nil, fmt.Errorf("failed to decoded response: %w", err)
	}
	return agents, nil
}

// obtener agente por id
func (ur *userRepository) GetAgentById(accessToken string, ownerID string) (dto.User, error) {
	url := fmt.Sprintf("%s/users/by-role/%s?role=AGENT", ur.userServiceURL, ownerID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return dto.User{}, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", accessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dto.User{}, fmt.Errorf("failed to connect to user service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
	}

	var user dto.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return dto.User{}, fmt.Errorf("failed to decoded response: %w", err)
	}
	return user, nil
}

// deshabilitar agente por id
func (ur *userRepository) DisableAgent(accessToken string, agentID string) error {
	url := fmt.Sprintf("%s/users/%s/by-role?role=AGENT", ur.userServiceURL, agentID)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request to user service: %w", err)
	}

	req.Header.Set("Authorization", accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to connect to user service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("user service returned unexpected status: %d", resp.StatusCode)
	}

	return nil
}

// deshabilitar cliente por id
func (ur *userRepository) DisableClient(accessToken string, clientID string) error {
	url := fmt.Sprintf("%s/users/%s/by-role?role=CLIENT", ur.userServiceURL, clientID)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request to user service: %w", err)
	}

	req.Header.Set("Authorization", accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to connect to user service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("user service returned unexpected status: %d", resp.StatusCode)
	}

	return nil
}

// deshabilitar propietario por id
func (ur *userRepository) DisableOwner(accessToken string, ownerID string) error {
	url := fmt.Sprintf("%s/users/%s/by-role?role=OWNER", ur.userServiceURL, ownerID)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request to user service: %w", err)
	}

	req.Header.Set("Authorization", accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to connect to user service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("user service returned unexpected status: %d", resp.StatusCode)
	}

	return nil
}

// habilitar agente por id
func (ur *userRepository) EnableAgent(accessToken string, agentID string) error {
	url := fmt.Sprintf("%s/users/%s/by-role/activate?role=AGENT", ur.userServiceURL, agentID)

	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request to user service: %w", err)
	}

	req.Header.Set("Authorization", accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to connect to user service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("user service returned unexpected status: %d", resp.StatusCode)
	}

	return nil
}

// habilitar cliente por id
func (ur *userRepository) EnableClient(accessToken string, clientID string) error {
	url := fmt.Sprintf("%s/users/%s/by-role/activate?role=CLIENT", ur.userServiceURL, clientID)

	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request to user service: %w", err)
	}

	req.Header.Set("Authorization", accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to connect to user service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("user service returned unexpected status: %d", resp.StatusCode)
	}

	return nil
}

// habilitar propietario por id
func (ur *userRepository) EnableOwner(accessToken string, ownerID string) error {
	url := fmt.Sprintf("%s/users/%s/by-role/activate?role=OWNER", ur.userServiceURL, ownerID)

	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request to user service: %w", err)
	}

	req.Header.Set("Authorization", accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to connect to user service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("user service returned unexpected status: %d", resp.StatusCode)
	}

	return nil
}
