package routes

import (
	"github.com/damshxy/api-car-go/database"
	"github.com/damshxy/api-car-go/handlers"
	"github.com/damshxy/api-car-go/repository"
	"github.com/damshxy/api-car-go/services"
	"github.com/damshxy/api-car-go/usecase"
	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Group) {
	auth := e.Group("/auth")

	userRepo := repository.NewUserRepository(database.DB)
	loggerService := services.NewLoggerService() // Assuming NewLoggerService creates a new instance with necessary configuration
	userUseCase := usecase.NewUserUsecase(userRepo)
	handlerAuth := handlers.NewAuthHandler(userUseCase, loggerService)
	
	auth.POST("/register", handlerAuth.Register)
	auth.POST("/login", handlerAuth.Login)
}