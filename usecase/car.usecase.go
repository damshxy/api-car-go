package usecase

import (
	"errors"

	dtos "github.com/damshxy/api-car-go/dto"
	"github.com/damshxy/api-car-go/helpers"
	"github.com/damshxy/api-car-go/models"
	"github.com/damshxy/api-car-go/repository"
)

type CarUsecase interface {
	GetAll(token string) ([]*dtos.CarResponse, error)
	GetById(id uint, token string) (*dtos.CarResponse, error)
	Create(req dtos.CarRequest, token string) (*dtos.CarResponse, error)
	Update(id uint, req dtos.CarRequest, token string) (*dtos.CarResponse, error)
	Delete(id uint, token string) error
}

type carUsecase struct {
	carRepository repository.CarRepository
}

func NewCarUsecase(carRepository repository.CarRepository) CarUsecase {
	return &carUsecase{
		carRepository: carRepository,
	}
}

func (c *carUsecase) GetAll(token string) ([]*dtos.CarResponse, error) {
	claims, err := helpers.ValidateJWT(token)
	if err != nil {
		return []*dtos.CarResponse{}, err
	}

	ownerID := claims["id"].(float64)

	cars, err := c.carRepository.GetAll()
	if err != nil {
		return []*dtos.CarResponse{}, err
	}

	carResponses := []*dtos.CarResponse {}
	for _, car := range cars {
		carResponse := dtos.CarResponse{
			ID: car.ID,
			NameCar: car.NameCar,
			PlateNumber: car.PlateNumber,
			OwnerID: int(ownerID),
		}
		carResponses = append(carResponses, &carResponse)
	}

	return carResponses, nil
}

func (c *carUsecase) GetById(id uint, token string) (*dtos.CarResponse, error) {
	claims, err := helpers.ValidateJWT(token)
	if err != nil {
		return &dtos.CarResponse{}, err
	}

	ownerID := claims["id"].(float64)

	car, err := c.carRepository.GetById(id)
	if err != nil {
		return &dtos.CarResponse{}, err
	}

	if car.OwnerID != int(ownerID) {
		return &dtos.CarResponse{}, errors.New("unauthorized")
	}

	carResponse := dtos.CarResponse{
		ID: car.ID,
		NameCar: car.NameCar,
		PlateNumber: car.PlateNumber,
		OwnerID: int(ownerID),
	}

	return &carResponse, nil
}

func (c *carUsecase) Create(req dtos.CarRequest, token string) (*dtos.CarResponse, error) {
	claims, err := helpers.ValidateJWT(token)
	if err != nil {
		return &dtos.CarResponse{}, err
	}

	ownerID := claims["id"].(float64)

	car := models.Car{
		NameCar: req.NameCar,
		PlateNumber: req.PlateNumber,
	}

	car.OwnerID = int(ownerID)
	if car.OwnerID != int(ownerID) {
		return &dtos.CarResponse{}, errors.New("unauthorized")
	}

	createdCar, err := c.carRepository.Create(&car)
	if err != nil {
		return &dtos.CarResponse{}, err
	}

	carResponse := dtos.CarResponse{
		ID: createdCar.ID,
		NameCar: createdCar.NameCar,
		PlateNumber: createdCar.PlateNumber,
		OwnerID: int(ownerID),
	}

	return &carResponse, nil
}

func (c *carUsecase) Update(id uint, req dtos.CarRequest, token string) (*dtos.CarResponse, error) {
	claims, err := helpers.ValidateJWT(token)
	if err != nil {
		return &dtos.CarResponse{}, err
	}

	ownerID := claims["id"].(float64)

	car, err := c.carRepository.GetById(id)
	if err != nil {
		return &dtos.CarResponse{}, err
	}

	car.NameCar = req.NameCar
	car.PlateNumber = req.PlateNumber

	updatedCar, err := c.carRepository.Update(car)
	if err != nil {
		return &dtos.CarResponse{}, err
	}

	carResponse := dtos.CarResponse {
		ID: updatedCar.ID,
		NameCar: updatedCar.NameCar,
		PlateNumber: updatedCar.PlateNumber,
		OwnerID: int(ownerID),
	}

	return &carResponse, nil
}

func (c *carUsecase) Delete(id uint, token string) error {
	claims, err := helpers.ValidateJWT(token)
	if err != nil {
		return err
	}

	ownerID := claims["id"].(float64)

	car, err := c.carRepository.GetById(id)
	if err != nil {
		return err
	}

	if car.OwnerID != int(ownerID) {
		return errors.New("unauthorized")
	}

	return c.carRepository.Delete(id)
}