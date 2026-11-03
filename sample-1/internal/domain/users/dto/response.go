package dto

import "time"

type UserResponse struct {
    ID        uint      `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Phone     string    `json:"phone"`
    Address   string    `json:"address"`
    Role      string    `json:"role"`
    Status    string    `json:"status"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type UserListResponse struct {
    Users      []UserResponse `json:"users"`
    Total      int64          `json:"total"`
    Page       int            `json:"page"`
    Limit      int            `json:"limit"`
    TotalPages int64          `json:"total_pages"`
}