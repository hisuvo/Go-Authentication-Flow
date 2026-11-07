package auth

import (
	"errors"
	"golang-authentication/internal/domain/user"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	// user opration
	CreateUser(user *user.User) error
	FindUserByEmail(email string) (*user.User, error)
	FindUserById(id uint) (*user.User, error)

	// Session operations.
	CreateSession(session *Session) error
	FindSessionByTokenHash(tokenHash string) (*Session, error)
	RevokeSession(sessionID uint) error
	RevokeAllUserSessions(userID uint) error

	// Cleanup.
	DeleteExpiredSessions() error
}

type repository struct {
	db *gorm.DB
}

func NewRespository (db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateUser(user *user.User) error {
	return r.db.Create(user).Error
}

// FindUserByEmail finds a user using their email address.
func (r *repository) FindUserByEmail(email string) (*user.User, error) {
	var user user.User

	err := r.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FindUserByID finds a user using their primary key.
func (r *repository) FindUserById(id uint) (*user.User, error) {
	var user user.User

	err := r.db.Where("id = ?", id).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// CreateSession stores a new authenticated session.
func (r *repository) CreateSession(session *Session) error {
	return r.db.Create(session).Error
}

// FindSessionByTokenHash finds a session using the hashed token.
func (r *repository) FindSessionByTokenHash(tokenHash string) (*Session, error) {

	var session Session

	err := r.db.
		Where("token_hash = ?", tokenHash).
		First(&session).
		Error

	if err != nil {
		return nil, err
	}

	return &session, nil
}

// RevokeSession invalidates one session.
func (r *repository) RevokeSession(sessionId uint) error {
	now := time.Now()

	result := r.db.Model(&Session{}).Where("sessionId = ?", sessionId).Where("revoked_at IS NULL").
		Update("revoked_at", &now)
	
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// RevokeAllUserSessions invalidates every active session
// belonging to a specific user.
func (r *repository) RevokeAllUserSessions(userID uint) error {
	now := time.Now()

	return r.db.
		Model(&Session{}).
		Where("user_id = ?", userID).
		Where("revoked_at IS NULL").
		Update("revoked_at", &now).
		Error
}

// DeleteExpiredSessions removes expired sessions.
//
// We also remove sessions that were revoked more than
// 30 days ago to keep the database clean.
func (r *repository) DeleteExpiredSessions() error {
	return r.db.
		Where(
			"expires_at < ? OR revoked_at < ?",
			time.Now(),
			time.Now().Add(-30*24*time.Hour),
		).
		Delete(&Session{}).
		Error
}

// Helper kept here if you want to explicitly check
// GORM's not-found error inside this package.
func isRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}