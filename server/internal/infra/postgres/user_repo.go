// Package postgres provides PostgreSQL implementations.
package postgres

import (
	"context"
	"errors"

	"github.com/arc-pub/server/internal/domain/auth"
	"github.com/arc-pub/server/internal/domain/user"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// UserRepo implements UserRepository with PostgreSQL.
type UserRepo struct {
	pool *pgxpool.Pool
}

// NewUserRepo creates a UserRepo with connection pool.
func NewUserRepo(pool *pgxpool.Pool) *UserRepo {
	return &UserRepo{pool: pool}
}

// FindByEmail retrieves user by email address.
func (r *UserRepo) FindByEmail(
	ctx context.Context,
	email string,
) (*user.User, error) {
	var u user.User
	var roleStr string

	err := r.pool.QueryRow(ctx, queryFindByEmail, email).Scan(
		&u.ID, &u.Email, &u.HashedPassword, &roleStr,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, auth.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	u.Role = user.Role(roleStr)
	return &u, nil
}

// Create inserts a new user into the database.
func (r *UserRepo) Create(ctx context.Context, u *user.User) error {
	_, err := r.pool.Exec(
		ctx, queryCreate, u.ID, u.Email, u.HashedPassword, u.Role,
	)
	return err
}

// ExistsByEmail checks if email is already registered.
func (r *UserRepo) ExistsByEmail(
	ctx context.Context,
	email string,
) (bool, error) {
	var exists bool
	err := r.pool.QueryRow(ctx, queryExistsByEmail, email).Scan(&exists)
	return exists, err
}
