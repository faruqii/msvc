package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/faruqii/msvc/authservice/internal/dto"
	"github.com/faruqii/msvc/authservice/internal/entities"
	"github.com/faruqii/msvc/authservice/internal/helper"
	"github.com/faruqii/msvc/authservice/internal/repository"
)

type AuthService interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (res *dto.RegisterResponse, err error)
	Login(ctx context.Context, req *dto.LoginRequest) (res *dto.LoginResponse, err error)
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) *authService {
	return &authService{userRepository: userRepository}
}

func (s *authService) Register(ctx context.Context, req *dto.RegisterRequest) (res *dto.RegisterResponse, err error) {
	existingUser, _ := s.userRepository.GetByEmail(req.Email)
	if existingUser.ID != "" {
		return nil, errors.New("user already exists")
	}

	hashedPwd, err := helper.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &entities.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPwd,
	}

	err = s.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return &dto.RegisterResponse{Message: "Registration successful"}, nil
}

func (s *authService) Login(ctx context.Context, req *dto.LoginRequest) (res *dto.LoginResponse, err error) {
	user, err := s.userRepository.GetByEmail(req.Email)
	if err != nil || user.ID == "" {
		return nil, fmt.Errorf("invalid email or password")
	}

	// Check password
	if !helper.CheckPasswordHash(req.Password, user.Password) {
		return nil, fmt.Errorf("invalid email or password")
	}

	// Generate JWT token (replace with your secret key and claims)
	tokenString, err := helper.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		Message: "Login successful",
		Token:   tokenString,
	}, nil
}
