// File: dto.go
// Purpose: Defines Data Transfer Objects for user profile and faction selection.
// ProfileResponse contains user data including faction and ships for the lobby.
// FactionRequest handles faction selection input. ShipDTO represents ship data
// for API responses. Keeps API contracts separate from domain entities.
// Path: server/internal/application/user/dto.go
// All Rights Reserved. Arc-Pub.

package user

// ProfileResponse is the output DTO for user profile.
type ProfileResponse struct {
	ID      string    `json:"id"`
	Email   string    `json:"email"`
	Faction *string   `json:"faction"`
	Ships   []ShipDTO `json:"ships"`
}

// ShipDTO represents a ship in API responses.
type ShipDTO struct {
	ID       string `json:"id"`
	ShipType string `json:"ship_type"`
	Slot     int    `json:"slot"`
}

// FactionRequest is the input DTO for faction selection.
type FactionRequest struct {
	Faction string `json:"faction"`
}

// FactionResponse confirms faction selection.
type FactionResponse struct {
	Faction  string  `json:"faction"`
	ShipID   string  `json:"ship_id"`
	ShipType string  `json:"ship_type"`
}
