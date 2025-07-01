package dto

type Status_Request struct {
	Status string `json:"status" binding:"required"`
}
