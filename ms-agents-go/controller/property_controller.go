package controller

import (
	"ms-agent/dto"
	"ms-agent/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PropertyController struct {
	service service.PropertyService
}

func NewPropertyController(service service.PropertyService) *PropertyController {
	return &PropertyController{service: service}
}

func (pc *PropertyController) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/properties", pc.ListProperties)
	rg.GET("/properties/:id", pc.GetPropertyByID)
	rg.GET("/properties/me", pc.GetMyProperties)
	rg.POST("/properties", pc.CreateProperty)
}

// obtener lista propiedades
func (pc *PropertyController) ListProperties(c *gin.Context) {
	properties, err := pc.service.ListProperties()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch properties"})
		return
	}

	c.JSON(http.StatusOK, properties)
}

// obtener propiedad por id
func (pc *PropertyController) GetPropertyByID(c *gin.Context) {
	id := c.Param("id")
	property, err := pc.service.GetPropertyByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Property not found"})
		return
	}
	c.JSON(http.StatusOK, property)
}

// obtener propiedades asignadas
func (pc *PropertyController) GetMyProperties(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		return
	}

	properties, err := pc.service.GetMyProperties(accessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, properties)
}

// crear nueva propiedad
func (pc *PropertyController) CreateProperty(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		return
	}

	var requestBody dto.PropertyRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	createdProperty, err := pc.service.CreateProperty(accessToken, requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdProperty)
}
