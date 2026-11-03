package dto

type CreateBlogRequest struct {
	Title       string `json:"title" validate:"required,min=3,max=255"`
	Description string `json:"description" validate:"required,min=10"`
}

type UpdateBlogRequest struct {
	Title       *string `json:"title" validate:"omitempty,min=3,max=255"`
	Description *string `json:"description" validate:"omitempty,min=10"`
}

type BlogQueryRequest struct {
	Search string `query:"search"`
	Page   int    `query:"page"`
	Limit  int    `query:"limit"`
	Sort   string `query:"sort"`
	Order  string `query:"order"`
}