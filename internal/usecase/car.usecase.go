package usecase

import (
	"errors"

	dtos "github.com/damshxy/api-car-go/internal/dto"
	"github.com/damshxy/api-car-go/internal/models"
	"github.com/damshxy/api-car-go/internal/repository"
	"github.com/damshxy/api-car-go/pkg/helpers"
	"github.com/damshxy/api-car-go/pkg/logger"
)

type CarUsecase interface {
	CreateCar(req *dtos.CarRequest, token string) (*dtos.CarResponse, error)
	FindAllCars(token string) ([]*dtos.CarResponse, error)
	FindCarByID(id uint, token string) (*dtos.CarResponse, error)
	UpdateCar(req *dtos.CarRequest, id uint, token string) (*dtos.CarResponse, error)
	DeleteCar(id uint, token string) error
}

type carUsecase struct {
	repo repository.CarRepository
	logger logger.LoggerService
}

func NewCarUsecase(repo repository.CarRepository) CarUsecase {
	return &carUsecase{
		repo: repo,
	}
}

func (u *carUsecase) CreateCar(req *dtos.CarRequest, token string) (*dtos.CarResponse, error) {
	claims, err := helpers.ValidateJWT(token)
	if err != nil {
		u.logger.Error("Failed to validate JWT" + err.Error())
		return nil, errors.New("failed to validate JWT")
	}

	userID := claims["id"].(float64)

	car := &models.Car{
		NameCar: req.NameCar,
		PlateNumber: req.PlateNumber,
		OwnerID: uint(userID),
	}

	createdCar, err := u.repo.Create(car)
	if err != nil {
		u.logger.Error("Failed to create car" + err.Error())
		return nil, errors.New("failed to create car")
	}

	carResponse := dtos.CarResponse{
		ID: createdCar.ID,
		NameCar: createdCar.NameCar,
		PlateNumber: createdCar.PlateNumber,
		OwnerID: uint(userID),
	}

	return &carResponse, nil
}

func (u *carUsecase) FindAllCars(token string) ([]*dtos.CarResponse, error) {
	claims, err := helpers.ValidateJWT(token)
	if err != nil {
		u.logger.Error("Failed to validate JWT" + err.Error())
		return nil, errors.New("failed to validate JWT")
	}

	userID := claims["id"].(float64)

	cars, err := u.repo.GetAll()
	if err != nil {
		u.logger.Error("Failed to get cars" + err.Error())
		return nil, errors.New("failed to get cars")
	}

	carResponses := []*dtos.CarResponse{}
	for _, car := range cars {
		if car.OwnerID == uint(userID) {
			carResponses = append(carResponses, &dtos.CarResponse{
				ID: car.ID,
				NameCar: car.NameCar,
				PlateNumber: car.PlateNumber,
				OwnerID: uint(userID),
			})
		}
	}

	return carResponses, nil
}

func (u *carUsecase) FindCarByID(id uint, token string) (*dtos.CarResponse, error) {
	claims, err := helpers.ValidateJWT(token)
	if err != nil {
		u.logger.Error("Failed to validate JWT: " + err.Error())
		return nil, errors.New("failed to validate JWT")
	}

	userID := claims["id"].(float64)

	car, err := u.repo.FindByID(id)
	if err != nil {
		u.logger.Error("Failed to get car: " + err.Error())
		return nil, errors.New("failed to get car")
	}

	if car.OwnerID != uint(userID) {
		u.logger.Error("Unauthorized access: user does not own the car")
		return nil, errors.New("unauthorized to access this car")
	}

	carResponse := &dtos.CarResponse{
		ID: car.ID,
		NameCar:     car.NameCar,
		PlateNumber: car.PlateNumber,
		OwnerID:     car.OwnerID,
	}

	return carResponse, nil
}


func (u *carUsecase) UpdateCar(req *dtos.CarRequest, id uint, token string) (*dtos.CarResponse, error) {
	claims, err := helpers.ValidateJWT(token)
	if err != nil {
		u.logger.Error("Failed to validate JWT" + err.Error())
		return nil, errors.New("failed to validate JWT")
	}

	userID := claims["id"].(float64)

	car, err := u.repo.FindByID(id)
	if err != nil {
		u.logger.Error("Failed to get car" + err.Error())
		return nil, errors.New("failed to get car")
	}

	if car.OwnerID != uint(userID) {
		u.logger.Error("Unauthorized to update this car")
		return nil, errors.New("unauthorized to update this car")
	}

	car.NameCar = req.NameCar
	car.PlateNumber = req.PlateNumber

	updatedCar, err := u.repo.Update(car)
	if err != nil {
		u.logger.Error("Failed to update car" + err.Error())
		return nil, errors.New("failed to update car")
	}

	carResponse := &dtos.CarResponse{
		ID: updatedCar.ID,
		NameCar: updatedCar.NameCar,
		PlateNumber: updatedCar.PlateNumber,
		OwnerID: uint(userID),
	}

	return carResponse, nil
}

func (u *carUsecase) DeleteCar(id uint, token string) error {
	claims, err := helpers.ValidateJWT(token)
	if err != nil {
		u.logger.Error("Failed to validate JWT" + err.Error())
		return errors.New("failed to validate JWT")
	}

	userID := claims["id"].(float64)

	car, err := u.repo.FindByID(id)
	if err != nil {
		u.logger.Error("Failed to get car" + err.Error())
		return errors.New("failed to get car")
	}

	if car.OwnerID != uint(userID) {
		u.logger.Error("Unauthorized to delete this car")
		return errors.New("unauthorized to delete this car")
	}

	return u.repo.Delete(car)
}