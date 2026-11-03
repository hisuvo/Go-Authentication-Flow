package blog

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

var (
    ErrBlogNotFound = errors.New("blog not found")
)

type Repository interface {
    Create(ctx context.Context, blog *Blog) error
    FindAll(
        ctx context.Context,
        search string,
        page int,
        limit int,
        sort string,
        order string,
    ) ([]Blog, int64, error)
    FindByID(ctx context.Context, id uint) (*Blog, error)
    Update(ctx context.Context, blog *Blog) error
    Delete(ctx context.Context, id uint) error
}

type repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
    return &repository{
        db: db,
    }
}

func (r *repository) Create(ctx context.Context, blog *Blog) error {
    return r.db.WithContext(ctx).Create(blog).Error
}

func (r *repository) FindAll(
    ctx context.Context,
    search string,
    page int,
    limit int,
    sort string,
    order string,
) ([]Blog, int64, error) {

    var blogs []Blog
    var total int64

    query := r.db.WithContext(ctx).Model(&Blog{})

    // Search by blog title or description.
    if search != "" {
        keyword := "%" + search + "%"

        query = query.Where(
            "title ILIKE ? OR description ILIKE ?",
            keyword,
            keyword,
        )
    }

    // Count records before applying pagination.
    if err := query.Count(&total).Error; err != nil {
        return nil, 0, err
    }

    // Whitelist sortable database columns.
    allowedSort := map[string]bool{
        "id":         true,
        "title":      true,
        "created_at": true,
        "updated_at": true,
    }

    if !allowedSort[sort] {
        sort = "created_at"
    }

    if order != "asc" && order != "desc" {
        order = "desc"
    }

    offset := (page - 1) * limit

    err := query.
        Order(sort + " " + order).
        Offset(offset).
        Limit(limit).
        Find(&blogs).Error

    if err != nil {
        return nil, 0, err
    }

    return blogs, total, nil
}

func (r *repository) FindByID(
    ctx context.Context,
    id uint,
) (*Blog, error) {

    var blog Blog

    err := r.db.WithContext(ctx).First(&blog, id).Error

    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, ErrBlogNotFound
    }

    if err != nil {
        return nil, err
    }

    return &blog, nil
}

func (r *repository) Update(ctx context.Context, blog *Blog) error {
    result := r.db.WithContext(ctx).
        Model(&Blog{}).
        Where("id = ?", blog.ID).
        Updates(map[string]interface{}{
            "title":       blog.Title,
            "description": blog.Description,
        })

    if result.Error != nil {
        return result.Error
    }

    if result.RowsAffected == 0 {
        return ErrBlogNotFound
    }

    return nil
}

func (r *repository) Delete(ctx context.Context, id uint) error {
    result := r.db.WithContext(ctx).Delete(&Blog{}, id)

    if result.Error != nil {
        return result.Error
    }

    if result.RowsAffected == 0 {
        return ErrBlogNotFound
    }

    return nil
}