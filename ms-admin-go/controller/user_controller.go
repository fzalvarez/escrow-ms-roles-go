package controller

import (
	"ms-admin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{service: service}
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	rg.GET("/clients", uc.ListClients)
	rg.GET("/clients/:id", uc.GetClientById)
	rg.GET("/owners", uc.ListOwners)
	rg.GET("/owners/:id", uc.GetOwnerById)
	rg.GET("/agents", uc.ListAgents)
	rg.GET("/agents/:id", uc.GetAgentById)
	rg.PUT("/agents/:id/status", uc.EnableAgent)
	rg.DELETE("/agents/:id/status", uc.DisableAgent)
	rg.PUT("/clients/:id/status", uc.EnableClient)
	rg.DELETE("/clients/:id/status", uc.DisableClient)
	rg.PUT("/owners/:id/status", uc.EnableOwner)
	rg.DELETE("/owners/:id/status", uc.DisableOwner)

}

// obtener lista clientes
func (uc *UserController) ListClients(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		return
	}
	clients, err := uc.service.GetClient(accessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch clients"})
		return
	}
	c.JSON(http.StatusOK, clients)
}

// obtener cliente por id
func (uc *UserController) GetClientById(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		return
	}

	id := c.Param("id")

	client, err := uc.service.GetClientById(accessToken, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch clients"})
		return
	}
	c.JSON(http.StatusOK, client)

}

// obtener lista propietarios
func (uc *UserController) ListOwners(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		return
	}
	owners, err := uc.service.GetOwner(accessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch clients"})
		return
	}
	c.JSON(http.StatusOK, owners)
}

// obtener propietario por id
func (uc *UserController) GetOwnerById(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		return
	}

	id := c.Param("id")

	client, err := uc.service.GetOwnerById(accessToken, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch clients"})
		return
	}
	c.JSON(http.StatusOK, client)

}

// obtener lista agentes
func (uc *UserController) ListAgents(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		return
	}
	agents, err := uc.service.GetAgent(accessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch clients"})
		return
	}
	c.JSON(http.StatusOK, agents)
}

// obtener agente por id
func (uc *UserController) GetAgentById(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		return
	}

	id := c.Param("id")

	client, err := uc.service.GetAgentById(accessToken, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch clients"})
		return
	}
	c.JSON(http.StatusOK, client)
}

// deshabilitar agent por id
func (uc *UserController) DisableAgent(c *gin.Context) {
	id := c.Param("id")

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		return
	}

	err := uc.service.DisableAgent(authHeader, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign agent"})
		return
	}

	c.Status(http.StatusNoContent)

}

// deshabilitar cliente por id
func (uc *UserController) DisableClient(c *gin.Context) {
	id := c.Param("id")

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		return
	}

	err := uc.service.DisableClient(authHeader, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign agent"})
		return
	}

	c.Status(http.StatusNoContent)
}

// deshabilitar propietario por id
func (uc *UserController) DisableOwner(c *gin.Context) {
	id := c.Param("id")

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		return
	}

	err := uc.service.DisableOwner(authHeader, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign agent"})
		return
	}

	c.Status(http.StatusNoContent)
}

// habilitar agente por id
func (uc *UserController) EnableAgent(c *gin.Context) {
	id := c.Param("id")

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		return
	}

	err := uc.service.EnableAgent(authHeader, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign agent"})
		return
	}

	c.Status(http.StatusNoContent)
}

// habilitar cliente por id
func (uc *UserController) EnableClient(c *gin.Context) {
	id := c.Param("id")

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		return
	}

	err := uc.service.EnableClient(authHeader, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign agent"})
		return
	}

	c.Status(http.StatusNoContent)
}

// habilitar propietario por id
func (uc *UserController) EnableOwner(c *gin.Context) {
	id := c.Param("id")

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
		return
	}

	err := uc.service.EnableOwner(authHeader, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign agent"})
		return
	}

	c.Status(http.StatusNoContent)
}
