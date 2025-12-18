// File: ship_repo.go
// Purpose: Implements ShipRepository interface using PostgreSQL with sqlc
// generated code. Provides Create for adding new ships, GetByUser to list
// all user ships, and Count for slot availability checks. Uses pgx pool for
// database access. Handles UUID and timestamp conversions between layers.
// Path: server/internal/infra/postgres/ship_repo.go
// All Rights Reserved. Arc-Pub.

package postgres

import (
	"context"

	domainShip "github.com/arc-pub/server/internal/domain/ship"
	"github.com/arc-pub/server/internal/infra/postgres/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ShipRepo implements ShipRepository.
type ShipRepo struct {
	queries *sqlc.Queries
}

// NewShipRepo creates a ShipRepo.
func NewShipRepo(pool *pgxpool.Pool) *ShipRepo {
	return &ShipRepo{queries: sqlc.New(pool)}
}

// Create inserts a new ship.
func (r *ShipRepo) Create(ctx context.Context, s *domainShip.Ship) error {
	return r.queries.CreateShip(ctx, sqlc.CreateShipParams{
		ID:       uuidToPgtype(s.ID),
		UserID:   uuidToPgtype(s.UserID),
		ShipType: s.ShipType,
		Slot:     int32(s.Slot),
	})
}

// GetByUser retrieves all ships for a user.
func (r *ShipRepo) GetByUser(
	ctx context.Context,
	userID uuid.UUID,
) ([]*domainShip.Ship, error) {
	rows, err := r.queries.GetUserShips(ctx, uuidToPgtype(userID))
	if err != nil {
		return nil, err
	}

	ships := make([]*domainShip.Ship, len(rows))
	for i, row := range rows {
		ships[i] = &domainShip.Ship{
			ID:       pgtypeToUUID(row.ID),
			UserID:   userID,
			ShipType: row.ShipType,
			Slot:     int(row.Slot),
		}
	}
	return ships, nil
}

// Count returns number of ships for a user.
func (r *ShipRepo) Count(ctx context.Context, userID uuid.UUID) (int, error) {
	count, err := r.queries.CountUserShips(ctx, uuidToPgtype(userID))
	return int(count), err
}
