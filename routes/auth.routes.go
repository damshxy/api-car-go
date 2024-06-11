package routes

import (
	"github.com/damshxy/api-car-go/database"
	"github.com/damshxy/api-car-go/handlers"
	"github.com/damshxy/api-car-go/repository"
	"github.com/damshxy/api-car-go/services"
	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Group) {
	auth := e.Group("/auth")

	userRepo := repository.NewUserRepository(database.DB)
	userService := services.NewUserServices(userRepo)
	handlerAuth := handlers.NewAuthHandler(userService)

	auth.POST("/register", handlerAuth.Register)
	auth.POST("/login", handlerAuth.Login)
}