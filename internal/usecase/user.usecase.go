package usecase

import (
	"errors"

	dtos "github.com/damshxy/api-car-go/internal/dto"
	"github.com/damshxy/api-car-go/internal/models"
	"github.com/damshxy/api-car-go/internal/repository"
	"github.com/damshxy/api-car-go/pkg/helpers"
	"github.com/damshxy/api-car-go/pkg/logger"
	"gorm.io/gorm"
)

type UserUsecase interface {
	Register(req *dtos.RegisterRequest) (*dtos.AuthResponse, error)
	Login(req *dtos.LoginRequest) (*dtos.AuthResponse, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
	logger logger.LoggerService
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
} 

func (u *userUsecase) Register(req *dtos.RegisterRequest) (*dtos.AuthResponse, error) {
	hashPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		u.logger.Error("Failed to hash password" + err.Error())
		return nil, errors.New("failed to hash password")
	}
	
	user := &models.User{
		Name: req.Name,
		Phone: req.Phone,
		Password: hashPassword,
	}

	createdUser, err := u.userRepo.Create(user)
	if err != nil {
		u.logger.Error("Failed to create user" + err.Error())
		return nil, errors.New("failed to create user")
	}

	authResponse := &dtos.AuthResponse{
		ID: int(createdUser.ID),
		Name: createdUser.Name,
		Phone: createdUser.Phone,
	}

	return authResponse, nil
}

func (u *userUsecase) Login(req *dtos.LoginRequest) (*dtos.AuthResponse, error) {
	user, err := u.userRepo.FindByPhone(req.Phone)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		u.logger.Error("Failed to find user" + err.Error())
		return nil, errors.New("failed to find user")
	}

	if user.Password == "" {
		u.logger.Error("user password is empty" + err.Error())
		return nil, errors.New("invalid credentials")
	}

	if err := helpers.ComparePassword(user.Password, req.Password); err != nil {
		u.logger.Error("Failed to compare password" + err.Error())
		return nil, errors.New("failed to compare password")
	}

	token, err := helpers.GenerateJWT(user.ID, user.Name)
	if err != nil {
		u.logger.Error("Failed to generate token" + err.Error())
		return nil, errors.New("failed to generate token")
	}

	authResponse := &dtos.AuthResponse{
		ID: int(user.ID),
		Name: user.Name,
		Phone: user.Phone,
		Token: token,
	}

	return authResponse, nil
}