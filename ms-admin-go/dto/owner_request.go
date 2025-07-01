package dto

type Owner_Request struct {
	OwnerID string `json:"ownerId" binding:"required"`
}
