package dto

import (
	"time"
)

type BlogResponse struct {
    ID          uint      `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type BlogListResponse struct {
    Blogs      []BlogResponse `json:"blogs"`
    Total      int64          `json:"total"`
    Page       int            `json:"page"`
    Limit      int            `json:"limit"`
    TotalPages int64          `json:"total_pages"`
}

