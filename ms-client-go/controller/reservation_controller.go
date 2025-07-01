package controller

import (
	"ms-clients/service"
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
