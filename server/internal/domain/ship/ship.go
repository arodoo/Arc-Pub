// File: ship.go
// Purpose: Defines the Ship entity representing a player's spacecraft. Each
// user can own up to 5 ships in slots 1-5. Ships have a type (e.g., betha_1)
// which determines their visual appearance and potentially stats. The initial
// ship type is determined by the user's faction for visual consistency.
// Path: server/internal/domain/ship/ship.go
// All Rights Reserved. Arc-Pub.

package ship

import "github.com/google/uuid"

// Ship represents a player's spacecraft.
type Ship struct {
	ID       uuid.UUID
	UserID   uuid.UUID
	ShipType string
	Slot     int
}

// NewShip creates a ship with validation.
func NewShip(userID uuid.UUID, shipType string, slot int) *Ship {
	return &Ship{
		ID:       uuid.New(),
		UserID:   userID,
		ShipType: shipType,
		Slot:     slot,
	}
}

// InitialShipType returns ship type for new users.
const InitialShipType = "betha_1"
