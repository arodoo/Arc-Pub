// File: faction.go
// Purpose: Defines the Faction type for user allegiance in the game. Each user
// belongs to exactly one faction (red, blue, or green) which determines their
// ship visual style and team alignment. Faction is immutable once selected.
// Provides constants and validation for the three supported faction values.
// Path: server/internal/domain/user/faction.go
// All Rights Reserved. Arc-Pub.

package user

// Faction represents user allegiance.
type Faction string

const (
	FactionRed   Faction = "red"
	FactionBlue  Faction = "blue"
	FactionGreen Faction = "green"
)

// IsValid checks if faction is a valid value.
func (f Faction) IsValid() bool {
	switch f {
	case FactionRed, FactionBlue, FactionGreen:
		return true
	}
	return false
}
