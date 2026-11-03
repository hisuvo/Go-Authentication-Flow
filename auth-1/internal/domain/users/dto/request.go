package dto

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Role     string `json:"role"`
}

type UpdateUserRequest struct {
	Name    string `json:"name" validate:"omitempty,min=2,max=100"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Role    string `json:"role"`
	Status  string `json:"status"`
}