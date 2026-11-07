package dto

import "time"

// UserResponse is the safe representation of a user returned by the API.
type UserResponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	Status     string `json:"status"`
}

// AuthResponse is returned after successful authentication.
type AuthResponse struct {
	AccessToken  string        `json:"accessToken"`
	RefreshToken string        `json:"refreshToken"`
	ExpiresAt    time.Time     `json:"expiresAt"`
	TokenType    string        `json:"tokenType"`
	User         UserResponse  `json:"user"`
}

// SessionResponse can be used when displaying session/device information.
type SessionResponse struct {
	ID         uint       `json:"id"`
	ExpiresAt  time.Time  `json:"expiresAt"`
	LastUsedAt *time.Time `json:"lastUsedAt,omitempty"`
	IPAddress  string     `json:"ipAddress,omitempty"`
	UserAgent  string     `json:"userAgent,omitempty"`
}