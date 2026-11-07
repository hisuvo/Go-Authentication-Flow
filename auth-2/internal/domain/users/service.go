package users

import "auth-two/internal/domain/users/dto"

type Service interface {
	Create(req *dto.UserCreateRequest) (*dto.UserResponse, error)
	Get()(*dto.UserResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Create(req *dto.UserCreateRequest) (*dto.UserResponse, error){
	return nil, nil
}

func (s *service) Get()(*dto.UserResponse, error){
	return nil, nil
}
