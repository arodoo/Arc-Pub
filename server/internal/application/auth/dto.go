// File: dto.go
// Purpose: Defines Data Transfer Objects for the authentication application
// layer. DTOs provide a clean interface between infrastructure and domain
// layers, preventing domain entities from leaking into HTTP responses. Includes
// LoginRequest for input validation, LoginResponse for API output with JSON
// tags following snake_case convention, and TokenPair for internal token data.
// Path: server/internal/application/auth/dto.go
// All Rights Reserved. Arc-Pub.

package auth

// LoginRequest is the input DTO for login.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse is the output DTO for login.
type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

// TokenPair holds access and refresh tokens.
type TokenPair struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int
}
