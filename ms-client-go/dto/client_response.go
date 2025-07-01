package dto

import "github.com/google/uuid"

type Profile struct {
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Country        string `json:"country"`
	ProfilePicture string `json:"profilePicture"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Cellphone string    `json:"cellphone"`
	Profile   Profile   `json:"profile"`
	Roles     []string  `json:"roles"`
}
