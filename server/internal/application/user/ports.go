// File: ports.go
// Purpose: Defines port interfaces for user application layer following
// hexagonal architecture. ProfileRepository handles user profile and faction
// persistence. ShipRepository manages ship CRUD operations. These abstractions
// enable testing with mocks and swapping implementations without changing logic.
// Path: server/internal/application/user/ports.go
// All Rights Reserved. Arc-Pub.

package user

import (
	"context"

	domainShip "github.com/arc-pub/server/internal/domain/ship"
	domainUser "github.com/arc-pub/server/internal/domain/user"
	"github.com/google/uuid"
)

// ProfileRepository defines user profile operations.
type ProfileRepository interface {
	GetProfile(ctx context.Context, userID uuid.UUID) (*Profile, error)
	SetFaction(ctx context.Context, userID uuid.UUID, f domainUser.Faction) error
}

// Profile holds user data for the profile endpoint.
type Profile struct {
	ID       uuid.UUID
	Email    string
	Faction  *domainUser.Faction
	ServerID *uuid.UUID
}

// ShipRepository defines ship persistence operations.
type ShipRepository interface {
	Create(ctx context.Context, s *domainShip.Ship) error
	GetByUser(ctx context.Context, userID uuid.UUID) ([]*domainShip.Ship, error)
	Count(ctx context.Context, userID uuid.UUID) (int, error)
}
