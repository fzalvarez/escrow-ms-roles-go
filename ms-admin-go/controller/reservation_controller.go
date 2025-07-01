package controller

import (
	"ms-admin/dto"
	"ms-admin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReservationController struct {
	service service.ReservationService
}

func NewReservationController(service service.ReservationService) *ReservationController {
	return &ReservationController{service: service}
}

func (rc *ReservationController) RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("/reservations/", rc.CreateRerservation)
	rg.PATCH("/reservations/:id/status/", rc.UpdateReservationStatus)
	rg.GET("/reservations/:id", rc.GetReservationByID)
	rg.GET("/reservations", rc.FindAll)
	rg.GET("/reservations/client/:id", rc.GetReservationsByClientId)
	rg.GET("/reservations/agent/:id", rc.GetReservationsByAgentId)

}

// crear reserva
func (rc *ReservationController) CreateRerservation(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		return
	}

	var reservation dto.ReservationRequest
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	createdReservation, err := rc.service.CreateRerservation(accessToken, &reservation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdReservation)
}

// actualizar status reserva
func (rc *ReservationController) UpdateReservationStatus(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	id := c.Param("id")

	var req dto.Status_Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	err := rc.service.UpdateStatus(accessToken, id, req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reservation status updated successfully"})
}

// obtener reserva por id
func (rc *ReservationController) GetReservationByID(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		return
	}

	reservationId := c.Param("id")

	reservation, err := rc.service.GetReservationById(accessToken, reservationId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
		return
	}
	c.JSON(http.StatusOK, reservation)
}

// obtener todas las reservas
func (rc *ReservationController) FindAll(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		return
	}
	reservations, err := rc.service.FindAll(accessToken)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
		return
	}
	c.JSON(http.StatusOK, reservations)
}

// obtener reservas por client id
func (rc *ReservationController) GetReservationsByClientId(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		return
	}

	clientId := c.Param("id")

	reservations, err := rc.service.GetReservationsByClientId(accessToken, clientId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
		return
	}
	c.JSON(http.StatusOK, reservations)
}

// obtener reservas por agent id
func (rc *ReservationController) GetReservationsByAgentId(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		return
	}

	agentId := c.Param("id")

	reservations, err := rc.service.GetReservationsByAgentId(accessToken, agentId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
		return
	}
	c.JSON(http.StatusOK, reservations)
}
