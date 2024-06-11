package handlers

import (
	"log"
	"net/http"
	"strconv"

	dtos "github.com/damshxy/api-car-go/dto"
	"github.com/damshxy/api-car-go/helpers"
	"github.com/damshxy/api-car-go/services"
	"github.com/labstack/echo/v4"
)

type CarHandler struct {
	CarService services.CarService
	Validator *helpers.CustomValidator
}

func NewCarHandler(carService services.CarService) *CarHandler {
	return &CarHandler{
		CarService: carService,
	}
}

func (h *CarHandler) GetAll(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")

	cars, err := h.CarService.GetAll(token)
	if err != nil {
		log.Printf("error in get all car: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get all cars",
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


	car, err := h.CarService.GetById(uint(id), token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get car by id",
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
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "required fields are missing",
		})
	}

	carCreated, err := h.CarService.Create(req, token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed to create car",
		})
	}

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
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "required fields are missing",
		})
	}

	carUpdated, err := h.CarService.Update(uint(id), req, token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed to update car",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update car",
		"car":     carUpdated,
	})
}

func (h *CarHandler) Delete(c echo.Context) error {
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

	err = h.CarService.Delete(uint(id), token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "failed to delete car",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete car",
	})
}