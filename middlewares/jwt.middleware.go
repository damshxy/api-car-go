package middlewares

import (
	"net/http"
	"strings"

	"github.com/damshxy/api-car-go/pkg/helpers"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "missing token",
			})
		}

		tokenString := strings.Split(authHeader, "Bearer ")[1]
		claims, err := helpers.ValidateJWT(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "invalid token",
			})
		}

		c.Set("user", claims)

		return next(c)
	}
}