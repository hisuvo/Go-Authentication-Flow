package user

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

var (
    ErrUserNotFound = errors.New("user not found")
    ErrEmailExists = errors.New("email already exists")
)

type Repository interface {
    Create(ctx context.Context, user *User) error
    FindAll(ctx context.Context) ([]User, error)
    FindByID(ctx context.Context, id uint) (*User, error)
    FindByEmail(ctx context.Context, email string) (*User, error)
    Update(ctx context.Context, user *User) error
    Delete(ctx context.Context, id uint) error
}

type repository struct {
    db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
    return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, user *User) error {
    return r.db.WithContext(ctx).Create(user).Error
}

func (r *repository) FindAll(ctx context.Context) ([]User, error) {
    var users []User

    err := r.db.WithContext(ctx).
        Order("created_at DESC").
        Find(&users).Error

    return users, err
}

func (r *repository) FindByID(ctx context.Context, id uint) (*User, error) {
    var user User

    err := r.db.WithContext(ctx).First(&user, id).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, ErrUserNotFound
    }

    if err != nil {
        return nil, err
    }

    return &user, nil
}

func (r *repository) FindByEmail(ctx context.Context, email string) (*User, error) {
    var user User

    err := r.db.WithContext(ctx).
        Where("email = ?", email).
        First(&user).Error

    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, ErrUserNotFound
    }

    if err != nil {
        return nil, err
    }

    return &user, nil
}

func (r *repository) Update(ctx context.Context, user *User) error {
    result := r.db.WithContext(ctx).
        Model(&User{}).
        Where("id = ?", user.ID).
        Updates(map[string]interface{}{
            "name":    user.Name,
            "phone":   user.Phone,
            "address": user.Address,
            "role":    user.Role,
            "status":  user.Status,
        })

    if result.Error != nil {
        return result.Error
    }

    if result.RowsAffected == 0 {
        return ErrUserNotFound
    }

    return nil
}

func (r *repository) Delete(ctx context.Context, id uint) error {
    result := r.db.WithContext(ctx).Delete(&User{}, id)

    if result.Error != nil {
        return result.Error
    }

    if result.RowsAffected == 0 {
        return ErrUserNotFound
    }

    return nil
}