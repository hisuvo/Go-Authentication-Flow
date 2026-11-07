package auth

import (
	"errors"
	"golang-authentication/internal/domain/user"
	"golang-authentication/internal/domain/user/dto"
	"strings"

	"gorm.io/gorm"
)

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
)

type Service interface {
	CreateUser(req *dto.RegisterRequest) (*dto.UserResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateUser(req *dto.RegisterRequest) (*dto.UserResponse, error) {

	email := strings.ToLower(strings.TrimSpace(req.Email))

	// Check email aready exist in database
	existingUser, err := s.repository.FindUserByEmail(email)

	if err != nil && existingUser != nil {
		return nil, ErrEmailAlreadyExists
	}

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound){
		return nil, err
	}

	// generate hash password here 

	user := &user.User{
		Name: req.Name,
		Email: email,
		Password: req.Password,
		Role: user.UserRole(req.Role),
		Status: user.UserStatus(req.Status),
	}


	if err := s.repository.CreateUser(user); err != nil {
		return nil, err
	}

	return nil, nil
}