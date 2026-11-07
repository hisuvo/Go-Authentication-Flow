package session

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
    gorm.Model

    UserID     uint `gorm:"not null;index"`
    TenantID   uint `gorm:"not null;index"`

    SessionKey string `gorm:"type:uuid;uniqueIndex;not null"`

    RefreshTokenHash string `gorm:"type:varchar(255);not null"`

    UserAgent string `gorm:"type:text"`
    IPAddress string `gorm:"type:varchar(45)"`

    ExpiresAt time.Time `gorm:"not null"`
    RevokedAt *time.Time
}