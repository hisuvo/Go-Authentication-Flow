package user

import (
	"authentication-test/internal/domain/users/dto"
	"context"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type Service interface {
    Create(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserResponse, error)
    FindAll(ctx context.Context) ([]dto.UserResponse, error)
    FindByID(ctx context.Context, id uint) (*dto.UserResponse, error)
    Update(ctx context.Context, id uint, req *dto.UpdateUserRequest) (*dto.UserResponse, error)
    Delete(ctx context.Context, id uint) error
}

type service struct {
    repository Repository
}

func NewService(repository Repository) Service {
    return &service{
        repository: repository,
    }
}

func (s *service) Create(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserResponse, error) {

    email := strings.ToLower(strings.TrimSpace(req.Email))

    _, err := s.repository.FindByEmail(ctx, email)
    if err == nil {
        return nil, ErrEmailExists
    }

    if !errors.Is(err, ErrUserNotFound) {
        return nil, err
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password),bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    role := strings.ToUpper(strings.TrimSpace(req.Role))
    if role == "" {
        role = "USER"
    }

    user := &User{
        Name:     strings.TrimSpace(req.Name),
        Email:    email,
        Password: string(hashedPassword),
        Phone:    req.Phone,
        Address:  req.Address,
        Role:     RoleStatus(role),
        Status:   "ACTIVE",
    }

    if err := s.repository.Create(ctx, user); err != nil {
        return nil, err
    }

    return toUserResponse(user), nil
}

func (s *service) FindAll(ctx context.Context) ([]dto.UserResponse, error) {
    users, err := s.repository.FindAll(ctx)
    if err != nil {
        return nil, err
    }

    responses := make([]dto.UserResponse, 0, len(users))

    for i := range users {
        responses = append(responses, *toUserResponse(&users[i]))
    }

    return responses, nil
}

func (s *service) FindByID(
    ctx context.Context,
    id uint,
) (*dto.UserResponse, error) {

    user, err := s.repository.FindByID(ctx, id)
    if err != nil {
        return nil, err
    }

    return toUserResponse(user), nil
}

func (s *service) Update(
    ctx context.Context,
    id uint,
    req *dto.UpdateUserRequest,
) (*dto.UserResponse, error) {

    user, err := s.repository.FindByID(ctx, id)
    if err != nil {
        return nil, err
    }

    if req.Name != "" {
        user.Name = strings.TrimSpace(req.Name)
    }

    if req.Phone != "" {
        user.Phone = req.Phone
    }

    if req.Address != "" {
        user.Address = req.Address
    }

    if req.Role != "" {
        role := string(strings.ToUpper(strings.TrimSpace(req.Role)))
        user.Role = RoleStatus(role)
    }

    if req.Status != "" {
        status := strings.ToUpper(strings.TrimSpace(req.Status))
        user.Status = UserStatus(status)
    }

    if err := s.repository.Update(ctx, user); err != nil {
        return nil, err
    }

    return toUserResponse(user), nil
}

func (s *service) Delete(ctx context.Context, id uint) error {
    return s.repository.Delete(ctx, id)
}

// Convert entity into API response.
// Password is intentionally excluded.
func toUserResponse(user *User) *dto.UserResponse {
    if user == nil {
        return nil
    }

    return &dto.UserResponse{
        ID:        user.ID,
        Name:      user.Name,
        Email:     user.Email,
        Phone:     user.Phone,
        Address:   user.Address,
        Role:      string(user.Role),
        Status:    string(user.Status),
        CreatedAt: user.CreatedAt,
        UpdatedAt: user.UpdatedAt,
    }
}

// Keep this import available only if you add database-specific
// error handling here in the future.
var _ = gorm.ErrRecordNotFound