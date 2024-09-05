package routes

import (
	"github.com/damshxy/api-car-go/config/database"
	"github.com/damshxy/api-car-go/internal/delivery/handlers"
	"github.com/damshxy/api-car-go/internal/repository"
	"github.com/damshxy/api-car-go/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

func authRoutes(app fiber.Router) {
	auth := app.Group("/auth")

	repo := repository.NewUserRepository(database.DB)
	usecase := usecase.NewUserUsecase(repo)
	handlers := handlers.NewUserHandler(usecase)

	auth.Post("/register", handlers.Register)
	auth.Post("/login", handlers.Login)
}