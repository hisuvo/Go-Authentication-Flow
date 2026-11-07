package user

import "golang-authentication/internal/domain/user/dto"

type Service interface {
	CreateUser(req *dto.RegisterRequest) (*dto.UserResponse, error)
}

type service struct {
	register Register
}

func NewService(register Register) Service {
	return &service{
		register: register,
	}
}

func (s *service) CreateUser(req *dto.RegisterRequest) (*dto.UserResponse, error){
	
	// Check email already exist in data base
	// hrer first hash password using sha-56

	user := &User{
		Name: req.Name,
		Email: req.Email,
		Password: req.Password,
		Role: UserRole(req.Role),
		Status: UserStatus(req.Status),
	}

	if err := s.register.CreateUser(user); err != nil {
		return nil, err
	}

	return ToUserResponse(user), nil
}