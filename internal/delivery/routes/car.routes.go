package routes

import (
	"github.com/damshxy/api-car-go/config/database"
	"github.com/damshxy/api-car-go/internal/delivery/handlers"
	"github.com/damshxy/api-car-go/internal/repository"
	"github.com/damshxy/api-car-go/internal/usecase"
	middlewares "github.com/damshxy/api-car-go/middleware"
	"github.com/gofiber/fiber/v2"
)

func carRoutes(app fiber.Router) {
	repo := repository.NewCarRepository(database.DB)
	usecase := usecase.NewCarUsecase(repo)
	handlers := handlers.NewCarHandler(usecase)

	app.Get("/cars", middlewares.JWTMiddleware, handlers.FindAllCars)
	app.Get("/car/:id", middlewares.JWTMiddleware, handlers.FindCarByID)
	app.Post("/car", middlewares.JWTMiddleware, handlers.CreateCar)
	app.Put("/car/:id", middlewares.JWTMiddleware, handlers.UpdateCar)
	app.Delete("/car/:id", middlewares.JWTMiddleware, handlers.DeleteCar)
}