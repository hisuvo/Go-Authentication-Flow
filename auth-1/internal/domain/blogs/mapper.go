package blog

import (
	"authentication-test/internal/domain/blogs/dto"
)

func ToBlogResponse(b *Blog) *dto.BlogResponse {
	if b == nil {
		return nil
	}

	return &dto.BlogResponse{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		CreatedAt:   b.CreatedAt,
		UpdatedAt:   b.UpdatedAt,
	}
}

func ToBlogListResponse(
	blogs []Blog,
	total int64,
	page int,
	limit int,
) *dto.BlogListResponse {

	responses := make([]dto.BlogResponse, 0, len(blogs))

	for i := range blogs {
		responses = append(responses, *ToBlogResponse(&blogs[i]))
	}

	totalPages := (total + int64(limit) - 1) / int64(limit)

	return &dto.BlogListResponse{
		Blogs:      responses,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}
}