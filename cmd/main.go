package main

import (
	"github.com/damshxy/api-car-go/config"
	"github.com/damshxy/api-car-go/database"
	"github.com/damshxy/api-car-go/internal/handlers/http/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.LoadConfig() 

	e := echo.New()

	// Database
	database.LoadDatabasePG(cfg)

	// Routes
	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":5050"))
}