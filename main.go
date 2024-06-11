package main

import (
	"log"

	"github.com/damshxy/api-car-go/database"
	"github.com/damshxy/api-car-go/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Database
	database.InitDB()

	// Routes
	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":5050"))
}