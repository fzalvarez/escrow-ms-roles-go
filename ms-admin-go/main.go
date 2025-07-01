package main

import (
	"log"
	"ms-admin/auth"
	"ms-admin/controller"
	"ms-admin/repository"
	"ms-admin/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	propertyRepo := repository.NewPropertyRepository("http://localhost:8082")
	propertyService := service.NewPropertyService(propertyRepo)
	propertyController := controller.NewPropertyController(propertyService)

	userRepo := repository.NewUserRepository("http://localhost:8080")
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	reservationRepo := repository.NewReservationRepository("http://localhost:8087")
	reservationService := service.NewReservationService(reservationRepo)
	reservationController := controller.NewReservationController(reservationService)

	r := gin.Default()

	api := r.Group("/admin")
	api.Use(auth.AuthMiddleware())
	api.Use(auth.RequireRoles("Role_SUPERADMIN"))
	propertyController.RegisterPropertyRoutes(api)
	userController.RegisterUserRoutes(api)
	reservationController.RegisterRoutes(api)

	r.Run()
}
