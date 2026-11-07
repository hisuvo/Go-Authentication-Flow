package dto

// RegisterRequest represents the request body
// used to create a new user.
type RegisterRequest struct {
	Name       string `json:"name" validate:"required,min=2,max=100"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=8,max=72"`
	Role       string `json:"role" validate:"required"`
	Status       string `json:"status"`
}

// LoginRequest represents the login request body.
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=72"`
}