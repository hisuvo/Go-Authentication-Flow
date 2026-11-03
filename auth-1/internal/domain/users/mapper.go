package user

import "authentication-test/internal/domain/users/dto"

func ToUserResponse(user *User) *dto.UserResponse {
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