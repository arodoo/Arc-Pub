// File: dto.go
// Purpose: Defines Data Transfer Objects for server listing and selection.
// ServerDTO contains server data for API responses. ServerSelectRequest
// handles server selection input. Keeps API contracts separate from domain.
// Path: server/internal/application/server/dto.go
// All Rights Reserved. Arc-Pub.

package server

// ServerDTO represents a server in API responses.
type ServerDTO struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Region string  `json:"region"`
	Host   *string `json:"host,omitempty"`
	Port   *int    `json:"port,omitempty"`
}

// ServerSelectRequest is the input for server selection.
type ServerSelectRequest struct {
	ServerID string `json:"server_id"`
}

// ServerSelectResponse confirms server selection.
type ServerSelectResponse struct {
	ServerID string `json:"server_id"`
	Name     string `json:"name"`
	Region   string `json:"region"`
}
