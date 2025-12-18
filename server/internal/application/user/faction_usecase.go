// File: faction_usecase.go
// Purpose: Implements faction selection use case for new users. Validates
// faction choice, sets it on the user (immutable once set), and creates the
// initial Betha 1 ship in slot 1. Returns error if user already has a faction.
// This is a one-time operation that determines user's team and ship style.
// Path: server/internal/application/user/faction_usecase.go
// All Rights Reserved. Arc-Pub.

package user

import (
	"context"
	"errors"

	domainShip "github.com/arc-pub/server/internal/domain/ship"
	domainUser "github.com/arc-pub/server/internal/domain/user"
	"github.com/google/uuid"
)

var ErrInvalidFaction = errors.New("invalid faction")
var ErrFactionAlreadySet = errors.New("faction already set")

// FactionUseCase handles faction selection.
type FactionUseCase struct {
	profiles ProfileRepository
	ships    ShipRepository
}

// NewFactionUseCase creates a FactionUseCase with dependencies.
func NewFactionUseCase(
	profiles ProfileRepository,
	ships ShipRepository,
) *FactionUseCase {
	return &FactionUseCase{profiles: profiles, ships: ships}
}

// Execute selects faction and creates initial ship.
func (uc *FactionUseCase) Execute(
	ctx context.Context,
	userID uuid.UUID,
	req FactionRequest,
) (*FactionResponse, error) {
	faction := domainUser.Faction(req.Faction)
	if !faction.IsValid() {
		return nil, ErrInvalidFaction
	}

	profile, err := uc.profiles.GetProfile(ctx, userID)
	if err != nil {
		return nil, err
	}

	if profile.Faction != nil {
		return nil, ErrFactionAlreadySet
	}

	if err := uc.profiles.SetFaction(ctx, userID, faction); err != nil {
		return nil, err
	}

	ship := domainShip.NewShip(userID, domainShip.InitialShipType, 1)
	if err := uc.ships.Create(ctx, ship); err != nil {
		return nil, err
	}

	return &FactionResponse{
		Faction:  string(faction),
		ShipID:   ship.ID.String(),
		ShipType: ship.ShipType,
	}, nil
}
