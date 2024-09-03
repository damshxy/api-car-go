package dtos

type CarRequest struct {
	NameCar string `json:"name_car" validate:"required"`
	PlateNumber string `json:"plate_number" validate:"required"`
}

type CarResponse struct {
	ID uint `json:"id"`
	NameCar string `json:"name_car"`
	PlateNumber string `json:"plate_number"`
	OwnerID uint `json:"owner_id"`
}