package dto

import (
	"time"

	"github.com/google/uuid"
)

type ReservationRequest struct {
	PropertyID uuid.UUID `json:"property_id"`
	ClientID   uuid.UUID `json:"client_id"`
	Amount     float64   `json:"amount"`
	FinalDate  time.Time `json:"final_date"`
}
