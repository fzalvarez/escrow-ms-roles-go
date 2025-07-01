package dto

import (
	"time"

	"github.com/google/uuid"
)

type ReservationResponse struct {
	ID         uuid.UUID `json:"ID"`
	PropertyID uuid.UUID `json:"property_id"`
	ClientID   uuid.UUID `json:"client_id"`
	AgentID    uuid.UUID `json:"agent_id"`
	Status     string    `json:"status"`
	Amount     float64   `json:"amount"`
	StartDate  time.Time `json:"start_date"`
	FinalDate  time.Time `json:"final_date"`
	CreatedAt  time.Time `json:"CreatedAt"`
	UpdatedAt  time.Time `json:"UpdatedAt"`
}
