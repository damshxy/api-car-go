package usecase

import (
	"errors"

	dtos "github.com/damshxy/api-car-go/dto"
	"github.com/damshxy/api-car-go/helpers"
	"github.com/damshxy/api-car-go/models"
	"github.com/damshxy/api-car-go/repository"
)

type UserUsecase interface {
	Register(req *dtos.RegisterRequest) (*models.User, error)
	Login(req *dtos.LoginRequest) (*dtos.AuthResponse, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
    return &userUsecase{
        userRepository: userRepository,
    }
}

func (u *userUsecase) Register(req *dtos.RegisterRequest) (*models.User, error) {
	hashPassword := helpers.HashingPassword([]byte(req.Password))

	user := &models.User{
		Name: req.Name,
		Phone: req.Phone,
		Password: string(hashPassword),
	}

	createdUser, err := u.userRepository.Create(user)
	if err != nil {
		return &models.User{}, err
	}

	return createdUser, nil
}

func (u *userUsecase) Login(req *dtos.LoginRequest) (*dtos.AuthResponse, error) {
	user, err := u.userRepository.FindByPhone(req.Phone)
    if err != nil {
        return &dtos.AuthResponse{}, err
    }
	hashedPassword := helpers.ComparePassword([]byte(user.Password), []byte(req.Password))
	if hashedPassword != nil {
		return &dtos.AuthResponse{}, errors.New("invalid credentials")
	}

	token, err := helpers.GenerateJWT(int(user.ID), user.Name, user.Phone)
	if err != nil {
		return &dtos.AuthResponse{}, err
	}

	response := &dtos.AuthResponse{
        ID:    int(user.ID),
        Name:  user.Name,
        Phone: user.Phone,
        Token: token,
    }

	return response, nil
}