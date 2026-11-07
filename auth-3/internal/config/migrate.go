package config

import (
	"golang-authentication/internal/domain/auth"
	"golang-authentication/internal/domain/user"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&user.User{},
		&auth.Session{},
	)
}