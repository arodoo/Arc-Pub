// File: ports.go
// Purpose: Defines port interfaces for server application layer following
// hexagonal architecture. ServerRepository handles server data persistence.
// UserServerRepository manages user-server assignment. These abstractions
// enable swapping implementations for real vs simulated servers.
// Path: server/internal/application/server/ports.go
// All Rights Reserved. Arc-Pub.

package server

import (
	"context"

	domainServer "github.com/arc-pub/server/internal/domain/server"
	"github.com/google/uuid"
)

// ServerRepository defines server data operations.
type ServerRepository interface {
	ListActive(ctx context.Context) ([]*domainServer.Server, error)
	GetByID(ctx context.Context, id uuid.UUID) (*domainServer.Server, error)
}

// UserServerRepository defines user-server assignment.
type UserServerRepository interface {
	SetServer(ctx context.Context, userID, serverID uuid.UUID) error
	GetUserServerID(ctx context.Context, userID uuid.UUID) (*uuid.UUID, error)
}
