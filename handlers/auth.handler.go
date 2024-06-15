package handlers

import (
	"net/http"

	dtos "github.com/damshxy/api-car-go/dto"
	"github.com/damshxy/api-car-go/helpers"
	"github.com/damshxy/api-car-go/services"
	"github.com/damshxy/api-car-go/usecase"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	UserUsecase   usecase.UserUsecase
	LoggerService services.LoggerService
	Validator     *helpers.CustomValidator
}

func NewAuthHandler(userUsecase usecase.UserUsecase, loggerService services.LoggerService) *AuthHandler {
	return &AuthHandler {
		UserUsecase: userUsecase,
		LoggerService: loggerService,
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

	user, err := h.UserUsecase.Register(&req)
	if err != nil {
		h.LoggerService.Error("Failed to register user: " + err.Error())
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	h.LoggerService.Info("User registered: " + user.Name)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success register user",
		"user": user,
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

	userLogin, err := h.UserUsecase.Login(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	h.LoggerService.Info("User logged in: " + userLogin.Name)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success login user",
		"user": userLogin,
	})
}