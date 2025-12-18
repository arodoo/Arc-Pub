// File: jwt_service.go
// Purpose: Implements TokenService interface using JSON Web Tokens (JWT) for
// stateless authentication. Generates access tokens (15 min expiry) for API
// authentication and refresh tokens (7 days) for session renewal. Uses HS256
// symmetric signing algorithm with configurable secret. Includes user ID and
// role in claims for authorization. Follows industry JWT best practices.
// Path: server/internal/infra/token/jwt_service.go
// All Rights Reserved. Arc-Pub.

package token

import (
	"time"

	"github.com/arc-pub/server/internal/application/auth"
	"github.com/arc-pub/server/internal/domain/user"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const (
	accessExpiry  = 15 * time.Minute
	refreshExpiry = 7 * 24 * time.Hour
)

// JWTService implements TokenService with JWT.
type JWTService struct {
	secret []byte
}

// NewJWTService creates a JWTService with secret.
func NewJWTService(secret string) *JWTService {
	return &JWTService{secret: []byte(secret)}
}

// GeneratePair creates access and refresh tokens.
func (s *JWTService) GeneratePair(
	userID uuid.UUID,
	role user.Role,
) (*auth.TokenPair, error) {
	access, err := s.generate(userID, role, accessExpiry)
	if err != nil {
		return nil, err
	}

	refresh, err := s.generate(userID, role, refreshExpiry)
	if err != nil {
		return nil, err
	}

	return &auth.TokenPair{
		AccessToken:  access,
		RefreshToken: refresh,
		ExpiresIn:    int(accessExpiry.Seconds()),
	}, nil
}

func (s *JWTService) generate(
	userID uuid.UUID,
	role user.Role,
	expiry time.Duration,
) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userID.String(),
		"role": string(role),
		"exp":  time.Now().Add(expiry).Unix(),
		"iat":  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}
