package user

import "golang-authentication/internal/domain/user/dto"

func ToUserResponse(user *User) *dto.UserResponse {
	return &dto.UserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Role: string(user.Role),
		Status: string(user.Status),
	}
}

func ToUserResponses (users []*User) []*dto.UserResponse {
	responses := make([]*dto.UserResponse, 0, len(users))

	for i := range users {
		responses = append(responses, ToUserResponse(users[i]))
	}

	return responses
}