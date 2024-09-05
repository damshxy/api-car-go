package models

type Car struct {
	ID uint `json:"id" gorm:"primary_key"`
	NameCar string `json:"name_car" gorm:"not null"`
	PlateNumber string `json:"plate_number" gorm:"not null; unique"`
	OwnerID uint `json:"owner_id" gorm:"not null"`
}