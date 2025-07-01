package main

import (
	"log"
	"ms-clients/auth"
	"ms-clients/controller"
	"ms-clients/repository"
	"ms-clients/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	userRepo := repository.NewUserRepository("http://localhost:8080")
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	reservationRepo := repository.NewReservationRepository("http://localhost:8087")
	reservationService := service.NewReservationService(reservationRepo)
	reservationController := controller.NewReservationController(reservationService)

	r := gin.Default()

	api := r.Group("/client")
	api.Use(auth.AuthMiddleware())
	api.Use(auth.RequireRoles("Role_CLIENT"))
	userController.RegisterRoutes(api)
	reservationController.RegisterAgentRoutes(api)

	r.Run()
}
