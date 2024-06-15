package dtos

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
	ID int `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Token string  `json:"token"`
}