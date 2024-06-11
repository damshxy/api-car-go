package routes

import "github.com/labstack/echo/v4"

func InitRoutes(e *echo.Echo) {
	apiGroup := e.Group("/api")

	AuthRoutes(apiGroup)
}