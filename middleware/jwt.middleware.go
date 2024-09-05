package middlewares

import (
	"strings"

	"github.com/damshxy/api-car-go/pkg/helpers"
	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	parts := strings.Split(authHeader, "Bearer ")
	if len(parts) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid authorization header format",
		})
	}

	tokenString := parts[1]

	claims, err := helpers.ValidateJWT(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid token",
		})
	}

	c.Locals("user", claims)

	return c.Next()
}