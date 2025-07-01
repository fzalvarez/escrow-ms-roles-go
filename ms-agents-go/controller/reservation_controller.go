package controller

import (
	"ms-agent/dto"
	"ms-agent/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReservationController struct {
	service service.ReservationService
}

func NewReservationController(service service.ReservationService) *ReservationController {
	return &ReservationController{service: service}
}

func (rc *ReservationController) RegisterAgentRoutes(rg *gin.RouterGroup) {
	rg.GET("/reservations/me", rc.GetMyReservations)
	rg.POST("/reservations/", rc.CreateRerservation)
	rg.PATCH("/reservations/:id/status/", rc.UpdateReservationStatus)
	rg.GET("/reservations/:id", rc.GetReservationByID)

}

// obtener mis reservas
func (rc *ReservationController) GetMyReservations(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
		return
	}

	reservations, err := rc.service.GetMyReservations(accessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reservations)
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

	var req dto.StatusRequest
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
