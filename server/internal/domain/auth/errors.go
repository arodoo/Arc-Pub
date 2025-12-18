// Arc-Pub - Metaverso 2D MMO Social
// Copyright (c) 2024. MIT License.

package auth

import "errors"

// Domain errors for authentication.
var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidToken       = errors.New("invalid token")
	ErrTokenExpired       = errors.New("token expired")
)
