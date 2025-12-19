// File: user_server_repo.go
// Purpose: Implements UserServerRepository for managing user-server
// assignments. Provides SetServer to assign server (one-time) and
// GetUserServerID to retrieve current assignment. Uses profile_repo
// queries for data access. Essential for server selection flow.
// Path: server/internal/infra/postgres/user_server_repo.go
// All Rights Reserved. Arc-Pub.

package postgres

import (
	"context"

	"github.com/arc-pub/server/internal/infra/postgres/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// UserServerRepo implements UserServerRepository.
type UserServerRepo struct {
	queries *sqlc.Queries
}

// NewUserServerRepo creates a UserServerRepo.
func NewUserServerRepo(pool *pgxpool.Pool) *UserServerRepo {
	return &UserServerRepo{queries: sqlc.New(pool)}
}

// SetServer assigns a server to user.
func (r *UserServerRepo) SetServer(
	ctx context.Context,
	userID, serverID uuid.UUID,
) error {
	return r.queries.SetUserServer(ctx, sqlc.SetUserServerParams{
		ServerID: uuidToPgtypeNullable(serverID),
		ID:       uuidToPgtype(userID),
	})
}

// GetUserServerID returns user's server ID if set.
func (r *UserServerRepo) GetUserServerID(
	ctx context.Context,
	userID uuid.UUID,
) (*uuid.UUID, error) {
	row, err := r.queries.GetUserProfile(ctx, uuidToPgtype(userID))
	if err != nil {
		return nil, err
	}
	if !row.ServerID.Valid {
		return nil, nil
	}
	id := pgtypeToUUID(row.ServerID)
	return &id, nil
}
