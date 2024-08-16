package handlers

import (
	"net/http"
	"strconv"

	dtos "github.com/damshxy/api-car-go/dto"
	"github.com/damshxy/api-car-go/helpers"
	"github.com/damshxy/api-car-go/services"
	"github.com/damshxy/api-car-go/usecase"
	"github.com/labstack/echo/v4"
)

type CarHandler struct {
	CarUsecase usecase.CarUsecase
	logger     services.LoggerService
	Validator *helpers.CustomValidator
}

func NewCarHandler(carUsecase usecase.CarUsecase, logger services.LoggerService) *CarHandler {
	return &CarHandler{
		CarUsecase: carUsecase,
		logger:     logger,
	}
}

func (h *CarHandler) GetAll(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "unauthorized",
		})
	}

	cars, err := h.CarUsecase.GetAll(token)
	if err != nil {
		h.logger.Error("Failed to get cars: " + err.Error())
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	h.logger.Info("Success get cars total " + strconv.Itoa(len(cars)))

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get cars",
		"cars":    cars,
	})
}

func (h *CarHandler) GetById(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "unauthorized",
		})
	}

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid id",
		})
	}


	car, err := h.CarUsecase.GetById(uint(id), token)
	if err != nil {
		h.logger.Error("Failed to get car by id: " + err.Error())
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	h.logger.Info("Success get car by id " + strconv.Itoa(int(car.ID)))
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get car",
		"car":     car,
	})
}

func (h *CarHandler) Create(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "unauthorized",
		})
	}

	var req dtos.CarRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid request",
		})
	}

	if err := h.Validator.Validate(req); err != nil {
		h.logger.Error("Required fields are missing: " + err.Error())
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "required fields are missing",
		})
	}

	carCreated, err := h.CarUsecase.Create(req, token)
	if err != nil {
		h.logger.Error("Failed to create car: " + err.Error())
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed to create car",
		})
	}

	h.logger.Info("Car created: " + carCreated.NameCar)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create car",
		"car":     carCreated,
	})
}

func (h *CarHandler) Update(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "unauthorized",
		})
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid id",
		})
	}

	var req dtos.CarRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid request",
		})
	}

	if err := h.Validator.Validate(req); err != nil {
		h.logger.Error("Required fields are missing: " + err.Error())
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "required fields are missing",
		})
	}

	carUpdated, err := h.CarUsecase.Update(uint(id), req, token)
	if err != nil {
		h.logger.Error("Failed to update car: " + err.Error())
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed to update car",
		})
	}

	h.logger.Info("Car updated: " + carUpdated.NameCar)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update car",
		"car":     carUpdated,
	})
}

func (h *CarHandler) Delete(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		h.logger.Error("Token is empty")
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "unauthorized",
		})
	}

	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Error("Invalid id: " + err.Error())
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid id",
		})
	}

	if err := h.CarUsecase.Delete(uint(id), token); err != nil {
		h.logger.Error("Failed to delete car: " + err.Error())
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	h.logger.Info("Car deleted by id: " + idStr)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete car",
	})
}