package repository

import (
	"github.com/damshxy/api-car-go/internal/models"
	"gorm.io/gorm"
)

type CarRepository interface {
	Create(car *models.Car) (*models.Car, error)
	FindByID(id uint) (*models.Car, error)
	Update(car *models.Car) (*models.Car, error)
	GetAll() ([]*models.Car, error)
	Delete(car *models.Car) error
}

type carRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) CarRepository {
	return &carRepository{
		db: db,
	}
}

func (r *carRepository) Create(car *models.Car) (*models.Car, error) {
	if err := r.db.Create(car).Error; err != nil {
		return nil, err
	}
	return car, nil
}


func (r *carRepository) FindByID(id uint) (*models.Car, error) {
	var car models.Car
	if err := r.db.First(&car, id).Error; err != nil {
		return nil, err
	}
	return &car, nil
}


func (r *carRepository) Update(car *models.Car) (*models.Car, error) {
	err := r.db.Save(&car).Error
	if err != nil {
		return nil, err
	}

	return car, nil
}

func (r *carRepository) GetAll() ([]*models.Car, error) {
	var cars []*models.Car
	err := r.db.Order("id").Find(&cars).Error
	if err != nil {
		return nil, err
	}

	return cars, nil
}

func (r *carRepository) Delete(car *models.Car) error {
	err := r.db.Delete(&car).Error
	if err != nil {
		return nil
	}

	return nil
}