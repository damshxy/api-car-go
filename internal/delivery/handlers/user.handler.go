package handlers

import (
	dtos "github.com/damshxy/api-car-go/internal/dto"
	"github.com/damshxy/api-car-go/internal/usecase"
	"github.com/damshxy/api-car-go/pkg/helpers"
	"github.com/damshxy/api-car-go/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
	logger logger.LoggerService
	validator *helpers.CustomValidator
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
		logger: logger.NewLoggerService(),
		validator: helpers.NewValidator(),
	}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req dtos.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		h.logger.Error("Failed to parse request body" + err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Failed to parse request body",
		})
	}

	if err := h.validator.Validate(req); err != nil {
		h.logger.Error("Failed to validate request body" + err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Failed to validate request body",
		})
	}

	authResponse, err := h.userUsecase.Register(&req)
	if err != nil {
		h.logger.Error("Failed to register user" + err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Failed to register user",
		})
	}

	h.logger.Info("User registered successfully")
	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"message": "User registered successfully",
		"data": authResponse,
	})
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req dtos.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		h.logger.Error("Failed to parse request body" + err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Failed to parse request body",
		})
	}

	if err := h.validator.Validate(req); err != nil {
		h.logger.Error("Failed to validate request body" + err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Failed to validate request body",
		})
	}

	authResponse, err := h.userUsecase.Login(&req)
	if err != nil {
		h.logger.Error("Failed to login user" + err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": "Failed to login user",
		})
	}

	h.logger.Info("User logged in successfully")
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "User logged in successfully",
		"data": authResponse,
	})
}