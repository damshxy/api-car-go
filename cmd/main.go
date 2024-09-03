package main

import (
	"github.com/damshxy/api-car-go/config"
	"github.com/damshxy/api-car-go/config/database"
	"github.com/damshxy/api-car-go/internal/delivery/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	cfg := config.LoadConfig()

	// Connect to database
	database.ConnectPostgres(cfg)

	// Initilize routes
	routes.InitilizeRoutes(app)

	// Start server
	app.Listen(":5050")
}