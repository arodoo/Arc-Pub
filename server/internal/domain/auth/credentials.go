// Arc-Pub - Metaverso 2D MMO Social
// Copyright (c) 2024. MIT License.

// Package auth contains authentication domain logic.
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
