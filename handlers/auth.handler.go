package handlers

import (
	"net/http"

	dtos "github.com/damshxy/api-car-go/dto"
	"github.com/damshxy/api-car-go/helpers"
	"github.com/damshxy/api-car-go/services"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	AuthService services.AuthService
	Validator *helpers.CustomValidator
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

func (h *AuthHandler) Register(c echo.Context) error {
	var req dtos.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid request",
		})
	}

	if err := h.Validator.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "required fields are missing",
		})
	}

	user, err := h.AuthService.Register(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success register user",
		"user":    user,
	})
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req dtos.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid request",
		})
	}

	if err := h.Validator.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "required fields are missing",
		})
	}

	userLogin, err := h.AuthService.Login(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success login user",
		"user":    userLogin,
	})
}