// File: login_usecase.go
// Purpose: Implements the login use case as the core application service for
// user authentication. Orchestrates the authentication flow: validates input
// credentials, retrieves user from repository, verifies password hash, and
// generates JWT token pair on success. Dependencies are injected via ports
// interfaces enabling testability and infrastructure independence as per DDD.
// Path: server/internal/application/auth/login_usecase.go
// All Rights Reserved. Arc-Pub.

package auth

import (
	"context"

	domainAuth "github.com/arc-pub/server/internal/domain/auth"
)

// LoginUseCase handles user authentication.
type LoginUseCase struct {
	users  UserRepository
	tokens TokenService
	hasher PasswordHasher
}

// NewLoginUseCase creates a LoginUseCase with dependencies.
func NewLoginUseCase(
	users UserRepository,
	tokens TokenService,
	hasher PasswordHasher,
) *LoginUseCase {
	return &LoginUseCase{
		users:  users,
		tokens: tokens,
		hasher: hasher,
	}
}

// Execute validates credentials and returns tokens.
func (uc *LoginUseCase) Execute(
	ctx context.Context,
	req LoginRequest,
) (*LoginResponse, error) {
	creds := domainAuth.NewCredentials(req.Email, req.Password)
	if !creds.IsValid() {
		return nil, domainAuth.ErrInvalidCredentials
	}

	usr, err := uc.users.FindByEmail(ctx, creds.Email)
	if err != nil {
		return nil, domainAuth.ErrInvalidCredentials
	}

	err = uc.hasher.Compare(usr.HashedPassword, creds.Password)
	if err != nil {
		return nil, domainAuth.ErrInvalidCredentials
	}

	pair, err := uc.tokens.GeneratePair(usr.ID, usr.Role)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken:  pair.AccessToken,
		RefreshToken: pair.RefreshToken,
		ExpiresIn:    pair.ExpiresIn,
	}, nil
}
