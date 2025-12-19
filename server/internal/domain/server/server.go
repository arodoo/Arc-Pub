// File: server.go
// Purpose: Defines the Server entity representing a game server instance.
// Contains server metadata including name, region, connection details (host/
// port) and active status. Null host indicates simulated server. Entity is
// used for server selection during user onboarding before faction choice.
// Path: server/internal/domain/server/server.go
// All Rights Reserved. Arc-Pub.

package server

import "github.com/google/uuid"

// Server represents a game server instance.
type Server struct {
	ID     uuid.UUID
	Name   string
	Region string
	Host   *string
	Port   *int
}

// IsSimulated returns true if server has no real connection.
func (s *Server) IsSimulated() bool {
	return s.Host == nil
}
