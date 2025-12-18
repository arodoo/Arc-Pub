// File: profile_usecase.go
// Purpose: Implements the profile use case for retrieving user data including
// faction and ships. Used by the lobby screen to display user state. Returns
// profile with null faction for new users triggering faction selection flow.
// Aggregates data from profile and ship repositories into a single response.
// Path: server/internal/application/user/profile_usecase.go
// All Rights Reserved. Arc-Pub.

package user

import (
	"context"

	"github.com/google/uuid"
)

// ProfileUseCase handles user profile retrieval.
type ProfileUseCase struct {
	profiles ProfileRepository
	ships    ShipRepository
}

// NewProfileUseCase creates a ProfileUseCase with dependencies.
func NewProfileUseCase(
	profiles ProfileRepository,
	ships ShipRepository,
) *ProfileUseCase {
	return &ProfileUseCase{profiles: profiles, ships: ships}
}

// Execute retrieves user profile with ships.
func (uc *ProfileUseCase) Execute(
	ctx context.Context,
	userID uuid.UUID,
) (*ProfileResponse, error) {
	profile, err := uc.profiles.GetProfile(ctx, userID)
	if err != nil {
		return nil, err
	}

	shipList, err := uc.ships.GetByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	ships := make([]ShipDTO, len(shipList))
	for i, s := range shipList {
		ships[i] = ShipDTO{
			ID:       s.ID.String(),
			ShipType: s.ShipType,
			Slot:     s.Slot,
		}
	}

	var faction *string
	if profile.Faction != nil {
		f := string(*profile.Faction)
		faction = &f
	}

	return &ProfileResponse{
		ID:      profile.ID.String(),
		Email:   profile.Email,
		Faction: faction,
		Ships:   ships,
	}, nil
}
