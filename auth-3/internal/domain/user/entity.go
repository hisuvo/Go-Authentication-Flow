package user

import "gorm.io/gorm"

type UserStatus string

const (
	UserStatusActive   UserStatus = "ACTIVE"
	UserStatusInactive UserStatus = "INACTIVE"
	UserStatusSuspended UserStatus = "SUSPENDED"
)

type UserRole string

const (
	UserRoleAdmin   UserRole = "ADMIN"
	UserRoleStaff   UserRole = "STAFF"
	UserRoleManager UserRole = "MANAGER"
)

type User struct {
	gorm.Model

	Name string `json:"name" gorm:"type:varchar(100);not null"`
	Email string `json:"email" gorm:"type:varchar(150);not null;uniqueIndex"`
	Password string `json:"-" gorm:"type:varchar(255);not null"`

	Role UserRole `json:"role" gorm:"type:varchar(30);not null;index"`
	Status UserStatus `json:"status" gorm:"type:varchar(30);not null;default:'ACTIVE';index"`
}
