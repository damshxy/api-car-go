package routes

import (
	"github.com/damshxy/api-car-go/database"
	handlers "github.com/damshxy/api-car-go/internal/handlers/http"
	"github.com/damshxy/api-car-go/internal/repository"
	"github.com/damshxy/api-car-go/internal/usecase"
	"github.com/damshxy/api-car-go/middlewares"
	"github.com/damshxy/api-car-go/pkg/logger"
	"github.com/labstack/echo/v4"
)

func CarRoutes(e *echo.Group) {
	carRepo := repository.NewCarRepository(database.DB)
	loggerService := logger.NewLoggerService()
	carUsecase := usecase.NewCarUsecase(carRepo)
	handlerCar := handlers.NewCarHandler(carUsecase, loggerService)

	e.GET("/cars", handlerCar.GetAll, middlewares.JWTMiddleware)
	e.GET("/car/:id", handlerCar.GetById, middlewares.JWTMiddleware)
	e.POST("/car", handlerCar.Create, middlewares.JWTMiddleware)
	e.PUT("/car/:id", handlerCar.Update, middlewares.JWTMiddleware)
	e.DELETE("/car/:id", handlerCar.Delete, middlewares.JWTMiddleware)
}