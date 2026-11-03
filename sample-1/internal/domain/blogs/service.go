package blog

import (
	"authentication-test/internal/domain/blogs/dto"
	"context"
	"errors"
	"strings"
)

type Service interface {
    Create(ctx context.Context, req *dto.CreateBlogRequest) (*dto.BlogResponse, error)
    FindAll(ctx context.Context, req *dto.BlogQueryRequest) (*dto.BlogListResponse, error)
    FindByID(ctx context.Context, id uint) (*dto.BlogResponse, error)
    Update(ctx context.Context, id uint, req *dto.UpdateBlogRequest) (*dto.BlogResponse, error)
    Delete(ctx context.Context, id uint) error
}

type service struct {
    repository Repository
}

func NewService(repository Repository) Service {
    return &service{
        repository: repository,
    }
}

func (s *service) Create(ctx context.Context,req *dto.CreateBlogRequest,) (*dto.BlogResponse, error) {

    title := strings.TrimSpace(req.Title)
    description := strings.TrimSpace(req.Description)

    if len(title) < 3 || len(title) > 255 {
        return nil, errors.New("title must contain 3 to 255 characters")
    }

    if len(description) < 10 {
        return nil, errors.New("description must contain at least 10 characters")
    }

    blog := &Blog{
        Title:       title,
        Description: description,
    }

    if err := s.repository.Create(ctx, blog); err != nil {
        return nil, err
    }

    return toBlogResponse(blog), nil
}

func (s *service) FindAll(ctx context.Context,req *dto.BlogQueryRequest) (*dto.BlogListResponse, error) {

    page := req.Page
    limit := req.Limit

    if page < 1 {
        page = 1
    }

    if limit < 1 {
        limit = 10
    }

    // Prevent excessively large database queries.
    if limit > 100 {
        limit = 100
    }

    sort := strings.ToLower(req.Sort)
    order := strings.ToLower(req.Order)

    blogs, total, err := s.repository.FindAll(
        ctx,
        strings.TrimSpace(req.Search),
        page,
        limit,
        sort,
        order,
    )

    if err != nil {
        return nil, err
    }

    responses := make([]dto.BlogResponse, 0, len(blogs))

    for i := range blogs {
        responses = append(responses, *toBlogResponse(&blogs[i]))
    }

    totalPages := (total + int64(limit) - 1) / int64(limit)

    return &dto.BlogListResponse{
        Blogs:      responses,
        Total:      total,
        Page:       page,
        Limit:      limit,
        TotalPages: totalPages,
    }, nil
}

func (s *service) FindByID(ctx context.Context,id uint,) (*dto.BlogResponse, error) {

    blog, err := s.repository.FindByID(ctx, id)
    if err != nil {
        return nil, err
    }

    return toBlogResponse(blog), nil
}

func (s *service) Update(ctx context.Context,id uint,req *dto.UpdateBlogRequest,) (*dto.BlogResponse, error) {

    blog, err := s.repository.FindByID(ctx, id)
    if err != nil {
        return nil, err
    }

    if req.Title != nil {
        title := strings.TrimSpace(*req.Title)

        if len(title) < 3 || len(title) > 255 {
            return nil, errors.New("title must contain 3 to 255 characters")
        }

        blog.Title = title
    }

    if req.Description != nil {
        description := strings.TrimSpace(*req.Description)

        if len(description) < 10 {
            return nil, errors.New("description must contain at least 10 characters")
        }

        blog.Description = description
    }

    if err := s.repository.Update(ctx, blog); err != nil {
        return nil, err
    }

    return toBlogResponse(blog), nil
}

func (s *service) Delete(ctx context.Context, id uint) error {
    return s.repository.Delete(ctx, id)
}

// Convert database entity into API response.
func toBlogResponse(blog *Blog) *dto.BlogResponse {
    if blog == nil {
        return nil
    }

    return &dto.BlogResponse{
        ID:          blog.ID,
        Title:       blog.Title,
        Description: blog.Description,
        CreatedAt:   blog.CreatedAt,
        UpdatedAt:   blog.UpdatedAt,
    }
}