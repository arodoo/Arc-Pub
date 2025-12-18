// File: errors.go
// Purpose: Defines domain-specific error types for the authentication bounded
// context. These sentinel errors provide clear, consistent error handling
// across the application. Using domain errors instead of generic ones enables
// proper error translation at the infrastructure layer (HTTP status codes)
// while keeping domain logic clean and infrastructure-agnostic as per DDD.
// Path: server/internal/domain/auth/errors.go
// All Rights Reserved. Arc-Pub.

package auth

import "errors"

// Domain errors for authentication.
var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidToken       = errors.New("invalid token")
	ErrTokenExpired       = errors.New("token expired")
)
