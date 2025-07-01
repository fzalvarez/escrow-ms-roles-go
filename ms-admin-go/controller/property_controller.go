package controller

import (
	"ms-admin/dto"
	"ms-admin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PropertyController struct {
	service service.PropertyService
}

func NewPropertyController(service service.PropertyService) *PropertyController {
	return &PropertyController{service: service}
}

func (pc *PropertyController) RegisterPropertyRoutes(rg *gin.RouterGroup) {
	rg.GET("/properties", pc.ListProperties)
	rg.GET("/properties/:id", pc.GetPropertyByID)
	rg.PUT("/properties/:id/update-owner", pc.UpdateOwner)
	rg.POST("/properties/:id/assign-agent", pc.AssignAgent)
	rg.PUT("properties/:id", pc.UpdateStatus)

}

// lista de propiedades
func (pc *PropertyController) ListProperties(c *gin.Context) {
	properties, err := pc.service.ListProperties()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch properties"})
		return
	}

	c.JSON(http.StatusOK, properties)
}

// propiedad por id
func (pc *PropertyController) GetPropertyByID(c *gin.Context) {
	id := c.Param("id")
	property, err := pc.service.GetPropertyByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Property not found"})
		return
	}
	c.JSON(http.StatusOK, property)
}

// actualizar propietario de una propiedad
func (pc *PropertyController) UpdateOwner(c *gin.Context) {
	id := c.Param("id")

	var req dto.Owner_Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		return
	}

	err := pc.service.UpdateOwner(authHeader, id, req.OwnerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update property owner"})
		return
	}

	c.Status(http.StatusNoContent)
}

// asignar agente a una propiedad
func (pc *PropertyController) AssignAgent(c *gin.Context) {
	id := c.Param("id")

	var req dto.Agent_Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		return
	}

	err := pc.service.AssignAgent(authHeader, id, req.AgentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign agent"})
		return
	}

	c.Status(http.StatusNoContent)
}

// actualizar status de una propiedad
func (pc *PropertyController) UpdateStatus(c *gin.Context) {
	id := c.Param("id")

	var req dto.Status_Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		return
	}

	err := pc.service.UpdateStatus(authHeader, id, req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign agent"})
		return
	}

	c.Status(http.StatusNoContent)

}
