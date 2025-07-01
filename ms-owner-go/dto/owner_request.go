package dto

type UpdateProfileRequest struct {
	FirstName      string `json:"firstName,omitempty"`
	LastName       string `json:"lastName,omitempty"`
	Country        string `json:"country,omitempty"`
	ProfilePicture string `json:"profilePicture,omitempty"`
}

type UpdateUserRequest struct {
	Email    string               `json:"email,omitempty"`
	Password string               `json:"password,omitempty"`
	Profile  UpdateProfileRequest `json:"profile,omitempty"`
}
