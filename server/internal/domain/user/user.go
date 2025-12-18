// File: user.go
// Purpose: Defines the User aggregate root for the identity bounded context.
// Contains core user identity data including UUID, email, hashed password, and
// role. As an aggregate root, User is the main entry point for user-related
// operations and ensures consistency. Includes Role type with predefined
// constants (user, admin) and helper methods like IsAdmin for authorization.
// Path: server/internal/domain/user/user.go
// All Rights Reserved. Arc-Pub.

package user

import "github.com/google/uuid"

// Role represents user authorization level.
type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

// User is the aggregate root for identity.
type User struct {
	ID             uuid.UUID
	Email          string
	HashedPassword string
	Role           Role
}

// NewUser creates a User with validated fields.
func NewUser(id uuid.UUID, email, hash string, role Role) *User {
	return &User{
		ID:             id,
		Email:          email,
		HashedPassword: hash,
		Role:           role,
	}
}

// IsAdmin checks if user has admin privileges.
func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}
