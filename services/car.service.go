package services

import (
	"errors"

	dtos "github.com/damshxy/api-car-go/dto"
	"github.com/damshxy/api-car-go/helpers"
	"github.com/damshxy/api-car-go/models"
	"github.com/damshxy/api-car-go/repository"
)

type CarService interface {
	GetAll(token string) ([]*dtos.CarResponse, error)
	GetById(id uint, token string) (*dtos.CarResponse, error)
	Create(req dtos.CarRequest, token string) (*dtos.CarResponse, error)
	Update(id uint, req dtos.CarRequest, token string) (*dtos.CarResponse, error)
	Delete(id uint, token string) error
}

type carService struct {
	repo repository.CarRepository
}

func NewCarServices(repo repository.CarRepository) CarService {
	return &carService{
		repo: repo,
	}
}

func (s *carService) GetAll(token string) ([]*dtos.CarResponse, error) {
	// Get the owner ID from the token
	claims, err := helpers.ValidateJWT(token)
	if err != nil {
		return []*dtos.CarResponse{}, err
	}

	// Convert the owner ID to an int
	ownerID := int(claims["id"].(float64))

	// Get all cars for the owner
	cars, err := s.repo.GetAll()
	if err != nil {
		return []*dtos.CarResponse{}, err
	}

	// Filter cars by owner ID
	carResponses := []*dtos.CarResponse{}
	for _, car := range cars {
		carResponse := dtos.CarResponse{
			ID: car.ID,
			NameCar: car.NameCar,
			PlateNumber: car.PlateNumber,
			OwnerID: ownerID,
		}
		carResponses = append(carResponses, &carResponse)
	}

	return carResponses, nil
}

func (s *carService) GetById(id uint, token string) (*dtos.CarResponse, error) {
	// Get the owner ID from the token
	claims, err := helpers.ValidateJWT(token)
	if err != nil {
		return &dtos.CarResponse{}, err
	}

	// Convert the owner ID to an int
	ownerID := int(claims["id"].(float64))

	// Get the car by ID
	car, err := s.repo.GetById(id)
	if err != nil {
		return &dtos.CarResponse{}, err
	}

	// Check if the car belongs to the owner
	if car.OwnerID != ownerID {
		return &dtos.CarResponse{}, errors.New("unauthorized")
	}

	// Create the car response
	carResponse := dtos.CarResponse{
		ID: car.ID,
		NameCar: car.NameCar,
		PlateNumber: car.PlateNumber,
		OwnerID: ownerID,
	}

	return &carResponse, nil
}

func (s *carService) Create(req dtos.CarRequest, token string) (*dtos.CarResponse, error) {
	// Get the owner ID from the token
	claims, err := helpers.ValidateJWT(token)
	if err != nil {
		return &dtos.CarResponse{}, err
	}

	// Convert the owner ID to an int
	ownerID := claims["id"].(float64)

	// Create the car
	car := models.Car{
		NameCar: req.NameCar,
		PlateNumber: req.PlateNumber,
	}

	// Set the owner ID
	car.OwnerID = int(ownerID)

	// save the car to the database
	createdCar, err := s.repo.Create(&car)
	if err != nil {
		return &dtos.CarResponse{}, err
	}

	// Create the car response
	carResponse := dtos.CarResponse{
		ID: createdCar.ID,
		NameCar: createdCar.NameCar,
		PlateNumber: createdCar.PlateNumber,
		OwnerID: int(ownerID),
	}

	return &carResponse, nil
}

func (s *carService) Update(id uint, req dtos.CarRequest, token string) (*dtos.CarResponse, error) {
	// Get the owner ID from the token
	claims, err :=helpers.ValidateJWT(token)
	if err != nil {
		return &dtos.CarResponse{}, err
	}

	// Convert the owner ID to an int
	ownerID := claims["id"].(float64)

	// Get the car by ID
	car, err := s.repo.GetById(id)
	if err != nil {
		return &dtos.CarResponse{}, err
	}

	car.NameCar = req.NameCar
	car.PlateNumber = req.PlateNumber

	// Check if the car belongs to the owner
	updatedCar, err := s.repo.Update(car)
	if err != nil {
		return &dtos.CarResponse{}, err
	}

	carResponse := dtos.CarResponse{
		ID: updatedCar.ID,
		NameCar: updatedCar.NameCar,
		PlateNumber: updatedCar.PlateNumber,
		OwnerID: int(ownerID),
	}

	return &carResponse, nil
}

func (s *carService) Delete(id uint, token string) error {
	claims, err := helpers.ValidateJWT(token)
	if err != nil {
		return err
	}

	ownerID := claims["id"].(float64)

	car, err := s.repo.GetById(id)
	if err != nil {
		return err
	}

	if car.OwnerID != int(ownerID) {
		return errors.New("unauthorized")
	}

	return s.repo.Delete(id)
}