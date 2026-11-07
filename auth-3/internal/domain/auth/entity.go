package auth

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model

	// UserID identifies the authenticated user who owns this session.
	UserID uint `json:"userId" gorm:"not null;index"`

	// Store only the SHA-256 hash of the session token. 
	// Never store the raw session token in the database.
	TokenHash string `json:"-" gorm:"type:char(64);not null;uniqueIndex"`

	// Session becomes invalid after this timestamp.
	ExpiresAt time.Time `json:"expiresAt" gorm:"not null;index"`

	// NULL means the session is currently active. 
	// A timestamp means the session has been revoked.
	RevokedAt *time.Time `json:"revokedAt"`

	// Useful for tracking session activity. 
	LastUsedAt *time.Time `json:"lastUsedAt"`

	// Optional security/audit information.
	IPAddress string `json:"ipAddress" gorm:"type:varchar(45)"`
	UserAgent string `json:"userAgent" gorm:"type:text"`
}
