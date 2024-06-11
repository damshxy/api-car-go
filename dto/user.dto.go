package dtos

import "github.com/damshxy/api-car-go/models"

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	User  *models.User `json:"user"`
	Token string  `json:"token"`
}