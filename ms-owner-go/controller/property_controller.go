package controller

import (
	"ms-owners/dto"
	"ms-owners/service"
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
	rg.POST("/properties", pc.CreateProperty)
	rg.GET("/properties/:id", pc.GetPropertyByID)
	rg.PUT("properties/:id", pc.UpdatePropertyByID)
	rg.DELETE("/properties/:id", pc.DeletePropertyByID)
	rg.GET("/properties", pc.GetMyProperties)

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

// obtener mis propiedades
func (pc *PropertyController) GetMyProperties(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
		return
	}

	properties, err := pc.service.GetMyProperties(authHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, properties)
}

// crear  propiedad
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

// actualizar propiedad
func (pc *PropertyController) UpdatePropertyByID(c *gin.Context) {
	id := c.Param("id")

	_, err := pc.service.GetPropertyByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Property not found"})
		return
	}

	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		return
	}

	var req dto.PropertyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	err = pc.service.UpdatePropertyByID(accessToken, id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "property updated successfully"})
}

// eliminar propiedad
func (pc *PropertyController) DeletePropertyByID(c *gin.Context) {
	id := c.Param("id")
	accessToken := c.GetHeader("Authorization")

	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing access token"})
		return
	}

	err := pc.service.DeletePropertyByID(accessToken, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Property deleted successfully"})
}
