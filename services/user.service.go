package services

import (
	"errors"

	dtos "github.com/damshxy/api-car-go/dto"
	"github.com/damshxy/api-car-go/helpers"
	"github.com/damshxy/api-car-go/models"
	"github.com/damshxy/api-car-go/repository"
)

type AuthService interface {
	Register(req dtos.RegisterRequest) (*models.User, error)
	Login(req dtos.LoginRequest) (*dtos.AuthResponse, error)
}

type authService struct {
	repo repository.UserRepository
}

func NewUserServices(repo repository.UserRepository) AuthService {
	return &authService{
		repo: repo,
	}
}

func (s *authService) Register(req dtos.RegisterRequest) (*models.User, error) {
	hashPassword := helpers.HashingPassword([]byte(req.Password))

	user := &models.User{
		Name: req.Name,
		Phone: req.Phone,
		Password: string(hashPassword),
	}

	createdUser, err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s *authService) Login(req dtos.LoginRequest) (*dtos.AuthResponse, error) {
	user, err := s.repo.FindByPhone(req.Phone)
	if err != nil {
		return &dtos.AuthResponse{}, err
	}

	hashPassword := helpers.ComparePassword([]byte(user.Password), []byte(req.Password))
	if hashPassword != nil {
		return &dtos.AuthResponse{}, errors.New("invalid credentials")
	}

	token, err := helpers.GenerateJWT(int(user.ID), user.Name, user.Phone)
	if err != nil {
		return &dtos.AuthResponse{}, err
	}

	
	return &dtos.AuthResponse{
		User:  user,
		Token: token,
	}, nil
}