package user

import "gorm.io/gorm"

type Register interface {
	CreateUser(user *User) error
}

type register struct {
	db *gorm.DB
}

func NewRegister(db *gorm.DB) Register {
	return &register{
		db: db,
	}
}

func (r *register) CreateUser(user *User) error {
	return r.db.Create(user).Error
}