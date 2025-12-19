// File: select_usecase.go
// Purpose: Implements server selection use case for new users. Validates
// server exists and is active, then assigns to user. Server selection is
// one-time only (immutable once set). Returns error if user already has
// server. This determines which game world the user joins permanently.
// Path: server/internal/application/server/select_usecase.go
// All Rights Reserved. Arc-Pub.

package server

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var ErrServerNotFound = errors.New("server not found")
var ErrServerAlreadySet = errors.New("server already set")

// SelectUseCase handles server selection.
type SelectUseCase struct {
	servers     ServerRepository
	userServers UserServerRepository
}

// NewSelectUseCase creates a SelectUseCase with dependencies.
func NewSelectUseCase(
	servers ServerRepository,
	userServers UserServerRepository,
) *SelectUseCase {
	return &SelectUseCase{servers: servers, userServers: userServers}
}

// Execute selects a server for the user.
func (uc *SelectUseCase) Execute(
	ctx context.Context,
	userID uuid.UUID,
	req ServerSelectRequest,
) (*ServerSelectResponse, error) {
	serverID, err := uuid.Parse(req.ServerID)
	if err != nil {
		return nil, ErrServerNotFound
	}

	existing, _ := uc.userServers.GetUserServerID(ctx, userID)
	if existing != nil {
		return nil, ErrServerAlreadySet
	}

	srv, err := uc.servers.GetByID(ctx, serverID)
	if err != nil {
		return nil, ErrServerNotFound
	}

	if err := uc.userServers.SetServer(ctx, userID, serverID); err != nil {
		return nil, err
	}

	return &ServerSelectResponse{
		ServerID: srv.ID.String(),
		Name:     srv.Name,
		Region:   srv.Region,
	}, nil
}
