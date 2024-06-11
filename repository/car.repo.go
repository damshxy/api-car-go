package repository

import (
	"github.com/damshxy/api-car-go/models"
	"gorm.io/gorm"
)

type CarRepository interface {
	Create(car *models.Car) (*models.Car, error)
	GetAll() ([]*models.Car, error)
	GetById(id uint) (*models.Car, error)
	Update(car *models.Car) (*models.Car, error)
	Delete(id uint) error
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

func (r *carRepository) GetAll() ([]*models.Car, error) {
	var cars []*models.Car
	
	if err := r.db.Find(&cars).Error; err != nil {
		return nil, err
	}

	return cars, nil
}

func (r *carRepository) GetById(id uint) (*models.Car, error) {
	var car models.Car
	if err := r.db.Where("id = ?", id).First(&car).Error; err != nil {
		return nil, err
	}

	return &car, nil
}

func (r *carRepository) Update(car *models.Car) (*models.Car, error) {
	if err := r.db.Save(car).Error; err != nil {
		return nil, err
	}

	return car, nil
}

func (r *carRepository) Delete(id uint) error {
	return r.db.Where("id = ?", id).Delete(&models.Car{}).Error
}