// Arc-Pub - Metaverso 2D MMO Social
// Copyright (c) 2024. MIT License.

// Package crypto provides password hashing.
package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 12

// BcryptHasher implements PasswordHasher with bcrypt.
type BcryptHasher struct{}

// NewBcryptHasher creates a BcryptHasher instance.
func NewBcryptHasher() *BcryptHasher {
	return &BcryptHasher{}
}

// Hash generates a bcrypt hash from password.
func (h *BcryptHasher) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcryptCost,
	)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Compare verifies password against hash.
func (h *BcryptHasher) Compare(hash, password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)
}
