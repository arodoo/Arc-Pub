package auth

import (
	"context"

	"github.com/arc-pub/server/internal/domain/user"
	"github.com/google/uuid"
)

// UserRepository defines user persistence operations.
type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*user.User, error)
	Create(ctx context.Context, u *user.User) error
	ExistsByEmail(ctx context.Context, email string) (bool, error)
}

// TokenService defines token generation operations.
type TokenService interface {
	GeneratePair(userID uuid.UUID, role user.Role) (*TokenPair, error)
}

// PasswordHasher defines password hashing operations.
type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hash, password string) error
}
