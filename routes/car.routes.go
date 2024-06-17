package routes

import (
	"github.com/damshxy/api-car-go/database"
	"github.com/damshxy/api-car-go/handlers"
	"github.com/damshxy/api-car-go/middlewares"
	"github.com/damshxy/api-car-go/repository"
	"github.com/damshxy/api-car-go/services"
	"github.com/damshxy/api-car-go/usecase"
	"github.com/labstack/echo/v4"
)

func CarRoutes(e *echo.Group) {
	carRepo := repository.NewCarRepository(database.DB)
	loggerService := services.NewLoggerService()
	carUsecase := usecase.NewCarUsecase(carRepo)
	handlerCar := handlers.NewCarHandler(carUsecase, loggerService)

	e.GET("/cars", handlerCar.GetAll, middlewares.JWTMiddleware)
	e.GET("/car/:id", handlerCar.GetById, middlewares.JWTMiddleware)
	e.POST("/car", handlerCar.Create, middlewares.JWTMiddleware)
	e.PUT("/car/:id", handlerCar.Update, middlewares.JWTMiddleware)
	e.DELETE("/car/:id", handlerCar.Delete, middlewares.JWTMiddleware)
}