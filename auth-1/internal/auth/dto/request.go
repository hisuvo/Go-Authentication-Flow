package dto

type LoginRequest struct {
	CoachingID uint   `json:"coaching_id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}