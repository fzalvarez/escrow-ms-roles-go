package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"ms-clients/dto"
	"net/http"
)

type UserRepository interface {
	GetMyProfile(accessToken string) (dto.User, error)
	UpdateProfile(accessToken string, profile dto.UpdateProfileRequest) error
}

type userRepository struct {
	userServiceURL string
}

func NewUserRepository(userServiceURL string) UserRepository {
	return &userRepository{userServiceURL: userServiceURL}
}

// obtener mi perfil
func (ur *userRepository) GetMyProfile(accessToken string) (dto.User, error) {
	url := fmt.Sprintf("%s/users/me", ur.userServiceURL)

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

	var userProfile dto.User
	if err := json.NewDecoder(resp.Body).Decode(&userProfile); err != nil {
		return dto.User{}, fmt.Errorf("fail to decode response: %w", err)
	}

	return userProfile, nil
}

// actualizar mi perfil
func (ur *userRepository) UpdateProfile(accessToken string, profile dto.UpdateProfileRequest) error {
	url := fmt.Sprintf("%s/users/profiles", ur.userServiceURL)
	body := dto.UpdateUserRequest{Profile: profile}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", accessToken)
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("→ [ms-agents] Status:", resp.StatusCode)
	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("→ [ms-agents] Response body:", string(respBody))

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("update failed, status=%d", resp.StatusCode)
	}
	return nil
}
