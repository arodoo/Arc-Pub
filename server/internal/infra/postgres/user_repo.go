// File: user_repo.go
// Purpose: Implements UserRepository interface using PostgreSQL with sqlc
// generated code. Provides data access methods for user persistence: find by
// email, create user, and check existence. Uses pgx connection pool for
// efficient database connections. Converts between sqlc-generated types and
// domain entities ensuring clean separation between persistence and domain.
// Path: server/internal/infra/postgres/user_repo.go
// All Rights Reserved. Arc-Pub.

package postgres

import (
	"context"
	"errors"

	"github.com/arc-pub/server/internal/domain/auth"
	"github.com/arc-pub/server/internal/domain/user"
	"github.com/arc-pub/server/internal/infra/postgres/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

// UserRepo implements UserRepository with sqlc.
type UserRepo struct {
	queries *sqlc.Queries
}

// NewUserRepo creates a UserRepo with connection pool.
func NewUserRepo(pool *pgxpool.Pool) *UserRepo {
	return &UserRepo{queries: sqlc.New(pool)}
}

// FindByEmail retrieves user by email address.
func (r *UserRepo) FindByEmail(
	ctx context.Context,
	email string,
) (*user.User, error) {
	row, err := r.queries.FindUserByEmail(ctx, email)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, auth.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return rowToUser(row), nil
}

// Create inserts a new user into the database.
func (r *UserRepo) Create(ctx context.Context, u *user.User) error {
	return r.queries.CreateUser(ctx, sqlc.CreateUserParams{
		ID:             uuidToPgtype(u.ID),
		Email:          u.Email,
		HashedPassword: u.HashedPassword,
		Role:           string(u.Role),
	})
}

// ExistsByEmail checks if email is already registered.
func (r *UserRepo) ExistsByEmail(
	ctx context.Context,
	email string,
) (bool, error) {
	return r.queries.ExistsUserByEmail(ctx, email)
}

func rowToUser(row sqlc.FindUserByEmailRow) *user.User {
	return &user.User{
		ID:             pgtypeToUUID(row.ID),
		Email:          row.Email,
		HashedPassword: row.HashedPassword,
		Role:           user.Role(row.Role),
	}
}

func uuidToPgtype(id uuid.UUID) pgtype.UUID {
	return pgtype.UUID{Bytes: id, Valid: true}
}

func uuidToPgtypeNullable(id uuid.UUID) pgtype.UUID {
	return pgtype.UUID{Bytes: id, Valid: true}
}

func pgtypeToUUID(id pgtype.UUID) uuid.UUID {
	return uuid.UUID(id.Bytes)
}
