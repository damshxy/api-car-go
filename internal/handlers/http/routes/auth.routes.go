package routes

import (
	"github.com/damshxy/api-car-go/database"
	handlers "github.com/damshxy/api-car-go/internal/handlers/http"
	"github.com/damshxy/api-car-go/internal/repository"
	"github.com/damshxy/api-car-go/internal/usecase"
	"github.com/damshxy/api-car-go/pkg/logger"
	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Group) {
	auth := e.Group("/auth")

	userRepo := repository.NewUserRepository(database.DB)
	loggerService := logger.NewLoggerService()
	userUseCase := usecase.NewUserUsecase(userRepo)
	handlerAuth := handlers.NewAuthHandler(userUseCase, loggerService)
	
	auth.POST("/register", handlerAuth.Register)
	auth.POST("/login", handlerAuth.Login)
}