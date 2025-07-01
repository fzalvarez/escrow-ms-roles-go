package dto

import (
	"time"

	"github.com/google/uuid"
)

type Location struct {
	Street  string  `json:"street"`
	City    string  `json:"city"`
	State   string  `json:"state"`
	Country string  `json:"country"`
	Zip     string  `json:"zip"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
}

type PropertyResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	OwnerID     uuid.UUID `json:"owner_id"`
	Location    Location  `json:"location"`
	Features    []string  `json:"features"`
	Images      []string  `json:"images"`
	CreatedAt   time.Time `json:"created_at"`
}
