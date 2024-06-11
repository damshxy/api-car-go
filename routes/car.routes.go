package routes

import (
	"github.com/damshxy/api-car-go/database"
	"github.com/damshxy/api-car-go/handlers"
	"github.com/damshxy/api-car-go/middlewares"
	"github.com/damshxy/api-car-go/repository"
	"github.com/damshxy/api-car-go/services"
	"github.com/labstack/echo/v4"
)

func CarRoutes(e *echo.Group) {
	carRepo := repository.NewCarRepository(database.DB)
	carService := services.NewCarServices(carRepo)
	handlerCar := handlers.NewCarHandler(carService)

	e.GET("/cars", handlerCar.GetAll, middlewares.JWTMiddleware)
	e.GET("/car/:id", handlerCar.GetById, middlewares.JWTMiddleware)
	e.POST("/car", handlerCar.Create, middlewares.JWTMiddleware)
	e.PUT("/car/:id", handlerCar.Update, middlewares.JWTMiddleware)
	e.DELETE("/car/:id", handlerCar.Delete, middlewares.JWTMiddleware)
}