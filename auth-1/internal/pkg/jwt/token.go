package jwt

import (
	"errors"
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   	uint 	`json:"user_id"`
	TenantID 	uint 	`json:"tenant_id"`
	Role 		string 	`json:"role"`
	SessionID  	string 	`json:"session_id"`
	TokenType  	string 	`json:"token_type"`

	jwtlib.RegisteredClaims
}

type Manager struct {
	Secret     []byte
	Issuer     string
	AccessTTL  time.Duration
}

func NewManager(secret string, issuer string, accessTTL time.Duration) *Manager {
	return &Manager{
		Secret: []byte(secret),
		Issuer: issuer,
		AccessTTL: accessTTL,
	}
}


func (m *Manager) GenerateAccessToken(userId uint, tenantId uint, role string, sessionId string,) (string, error) {
	now := time.Now()

	claim := Claims{
		UserID: userId,
		TenantID: tenantId,
		SessionID: sessionId,
		Role: role,
		TokenType: "access",
		RegisteredClaims: jwtlib.RegisteredClaims{
			Subject: "user",
			Issuer: m.Issuer,
			IssuedAt: jwtlib.NewNumericDate(now),
			ExpiresAt: jwtlib.NewNumericDate(now.Add(m.AccessTTL)),
		},
	}

	token := jwtlib.NewWithClaims(jwtlib.SigningMethodES256, claim)

	return token.SignedString(m.Secret)
}

func (m *Manager) ValidateAccessToken(tokenString string) (*Claims, error) {

	token, err := jwtlib.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwtlib.Token) (any, error) {
			if token.Method != jwtlib.SigningMethodHS256 {
				return nil, errors.New("unexpected signing method")
			}

			return m.Secret, nil
		},
		jwtlib.WithIssuer(m.Issuer),
		jwtlib.WithExpirationRequired(),
		jwtlib.WithValidMethods([]string{"HS256"}),
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	if claims.TokenType != "access" {
		return nil, errors.New("invalid token type")
	}

	return claims, nil
}
