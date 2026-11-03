package user

import (
	"gorm.io/gorm"
)

type RoleStatus string

const (
	RoleSuperAdmin RoleStatus = "SUPER_ADMIN"
	RoleAdmin      RoleStatus = "ADMIN"
	RoleManager    RoleStatus = "MANAGER"
	RoleAccountant RoleStatus = "ACCOUNTANT"
	RoleAuthor    RoleStatus = "AUTHOR"
	RoleReader    RoleStatus = "READER"
)

type UserStatus string

const (
	UserActive   UserStatus = "ACTIVE"
	UserInactive UserStatus = "INACTIVE"
	UserBlocked  UserStatus = "BLOCKED"
)


type User struct {
    gorm.Model

    Name     string `gorm:"type:varchar(100);not null"`
    Email    string `gorm:"type:varchar(150);uniqueIndex;not null"`
    Password string `gorm:"type:varchar(255);not null"`
    Phone    string `gorm:"type:varchar(20)"`
    Address  string `gorm:"type:text"`
    Role     RoleStatus `gorm:"type:varchar(20);not null;default:USER"`
    Status   UserStatus `gorm:"type:varchar(20);not null;default:ACTIVE"`
}