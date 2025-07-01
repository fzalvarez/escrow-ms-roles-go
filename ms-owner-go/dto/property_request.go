package dto

import "github.com/google/uuid"

type PropertyRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	OwnerID     uuid.UUID `json:"owner_id"`
	Location    Location  `json:"location"`
}
