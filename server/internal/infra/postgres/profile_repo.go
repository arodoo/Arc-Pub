// File: profile_repo.go
// Purpose: Implements ProfileRepository interface using PostgreSQL with sqlc
// generated code. Provides GetProfile to retrieve user profile data including
// faction and server. SetFaction updates user's faction choice. Uses pgx pool
// for database access. Converts between sqlc types and domain entities.
// Path: server/internal/infra/postgres/profile_repo.go
// All Rights Reserved. Arc-Pub.

package postgres

import (
	"context"

	appUser "github.com/arc-pub/server/internal/application/user"
	domainUser "github.com/arc-pub/server/internal/domain/user"
	"github.com/arc-pub/server/internal/infra/postgres/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ProfileRepo implements ProfileRepository.
type ProfileRepo struct {
	queries *sqlc.Queries
}

// NewProfileRepo creates a ProfileRepo.
func NewProfileRepo(pool *pgxpool.Pool) *ProfileRepo {
	return &ProfileRepo{queries: sqlc.New(pool)}
}

// GetProfile retrieves user profile data.
func (r *ProfileRepo) GetProfile(
	ctx context.Context,
	userID uuid.UUID,
) (*appUser.Profile, error) {
	row, err := r.queries.GetUserProfile(ctx, uuidToPgtype(userID))
	if err != nil {
		return nil, err
	}

	var faction *domainUser.Faction
	if row.Faction.Valid {
		f := domainUser.Faction(row.Faction.FactionType)
		faction = &f
	}

	var serverID *uuid.UUID
	if row.ServerID.Valid {
		id := pgtypeToUUID(row.ServerID)
		serverID = &id
	}

	return &appUser.Profile{
		ID:       pgtypeToUUID(row.ID),
		Email:    row.Email,
		Faction:  faction,
		ServerID: serverID,
	}, nil
}

// SetFaction updates user's faction.
func (r *ProfileRepo) SetFaction(
	ctx context.Context,
	userID uuid.UUID,
	f domainUser.Faction,
) error {
	return r.queries.SetUserFaction(ctx, sqlc.SetUserFactionParams{
		Faction: sqlc.NullFactionType{FactionType: sqlc.FactionType(f), Valid: true},
		ID:      uuidToPgtype(userID),
	})
}
