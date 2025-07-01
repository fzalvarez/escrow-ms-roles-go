package main

import (
	"log"
	"ms-owners/auth"
	"ms-owners/controller"
	"ms-owners/repository"
	"ms-owners/service"

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

	propertyRepo := repository.NewPropertyRepository("http://localhost:8082")
	propertyService := service.NewPropertyService(propertyRepo)
	propertyController := controller.NewPropertyController(propertyService)

	r := gin.Default()

	api := r.Group("/owner")
	api.Use(auth.AuthMiddleware())
	api.Use(auth.RequireRoles("Role_OWNER"))
	userController.RegisterRoutes(api)
	propertyController.RegisterRoutes(api)

	r.Run()
}
