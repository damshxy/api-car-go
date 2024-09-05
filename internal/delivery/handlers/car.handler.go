package handlers

import (
	"strconv"

	dtos "github.com/damshxy/api-car-go/internal/dto"
	"github.com/damshxy/api-car-go/internal/usecase"
	"github.com/damshxy/api-car-go/pkg/helpers"
	"github.com/damshxy/api-car-go/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type CarHandler struct {
	carUsecase usecase.CarUsecase
	logger logger.LoggerService
	validator *helpers.CustomValidator
}

func NewCarHandler(carUsecase usecase.CarUsecase) *CarHandler {
	return &CarHandler{
		carUsecase: carUsecase,
		logger: logger.NewLoggerService(),
		validator: helpers.NewValidator(),
	}
}

func (h *CarHandler) CreateCar(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		h.logger.Error("Failed to get token" + token)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	var req dtos.CarRequest
	if err := c.BodyParser(&req); err != nil {
		h.logger.Error("Failed to parse body" + err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	if err := h.validator.Validate(req); err != nil {
		h.logger.Error("Failed to validate body" + err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	carResponse, err := h.carUsecase.CreateCar(&req, token)
	if err != nil {
		h.logger.Error("Failed to create car" + err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create car",
		})
	}

	h.logger.Info("Car created successfully")
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Car created successfully",
		"data": carResponse,
	})
}

func (h *CarHandler) FindAllCars(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		h.logger.Error("Failed to get token" + token)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	carResponses, err := h.carUsecase.FindAllCars(token)
	if err != nil {
		h.logger.Error("Failed to get cars" + err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get cars",
		})
	}

	h.logger.Info("Cars retrieved successfully")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Cars retrieved successfully",
		"data": carResponses,
	})
}

func (h *CarHandler) FindCarByID(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		h.logger.Error("Failed to get token" + token)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Error("Failed to convert id" + err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	carResponse, err := h.carUsecase.FindCarByID(uint(id), token)
	if err != nil {
		h.logger.Error("Failed to get car" + err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get car",
		})
	}

	h.logger.Info("Car retrieved successfully")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Car retrieved successfully",
		"data": carResponse,
	})
}

func (h *CarHandler) UpdateCar(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		h.logger.Error("Failed to get token" + token)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Error("Failed to convert id" + err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	var req dtos.CarRequest
	if err := c.BodyParser(&req); err != nil {
		h.logger.Error("Failed to parse body" + err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	if err := h.validator.Validate(req); err != nil {
		h.logger.Error("Failed to validate body" + err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	carResponse, err := h.carUsecase.UpdateCar(&req, uint(id), token)
	if err != nil {
		h.logger.Error("Failed to update car" + err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update car",
		})
	}

	h.logger.Info("Car updated successfully")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Car updated successfully",
		"data": carResponse,
	})
}

func (h *CarHandler) DeleteCar(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		h.logger.Error("Failed to get token" + token)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Error("Failed to convert id" + err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	err = h.carUsecase.DeleteCar(uint(id), token)
	if err != nil {
		h.logger.Error("Failed to delete car" + err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete car",
		})
	}

	h.logger.Info("Car deleted successfully")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Car deleted successfully",
	})
}