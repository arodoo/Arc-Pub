// File: credentials.go
// Purpose: Defines the Credentials value object for authentication input.
// Encapsulates email and password as an immutable pair with validation logic.
// Value objects are compared by their attributes rather than identity, making
// this ideal for login request data. Provides IsValid method to check required
// fields before processing authentication attempts in the domain layer.
// Path: server/internal/domain/auth/credentials.go
// All Rights Reserved. Arc-Pub.

package auth

// Credentials is a value object for login input.
type Credentials struct {
	Email    string
	Password string
}

// NewCredentials creates validated credentials.
func NewCredentials(email, password string) Credentials {
	return Credentials{
		Email:    email,
		Password: password,
	}
}

// IsValid checks if credentials have required fields.
func (c Credentials) IsValid() bool {
	return c.Email != "" && c.Password != ""
}
