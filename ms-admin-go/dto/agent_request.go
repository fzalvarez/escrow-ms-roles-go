package dto

type Agent_Request struct {
	AgentID string `json:"agent_id" binding:"required"`
}
