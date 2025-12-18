// File: bcrypt_hasher.go
// Purpose: Implements the PasswordHasher interface using bcrypt algorithm.
// Bcrypt is the industry standard for password hashing with built-in salting
// and configurable cost factor. Uses cost 12 providing strong security while
// maintaining acceptable performance. Hash method generates secure hashes,
// Compare method safely verifies passwords against stored hashes using
// constant-time comparison to prevent timing attacks.
// Path: server/internal/infra/crypto/bcrypt_hasher.go
// All Rights Reserved. Arc-Pub.

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
