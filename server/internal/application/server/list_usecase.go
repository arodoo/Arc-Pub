// File: list_usecase.go
// Purpose: Implements use case for listing available game servers. Returns
// all active servers for display in server selection UI. Converts domain
// entities to DTOs for API response format. Used during user onboarding
// before faction selection to choose game server.
// Path: server/internal/application/server/list_usecase.go
// All Rights Reserved. Arc-Pub.

package server

import "context"

// ListUseCase handles server listing.
type ListUseCase struct {
	servers ServerRepository
}

// NewListUseCase creates a ListUseCase with dependencies.
func NewListUseCase(servers ServerRepository) *ListUseCase {
	return &ListUseCase{servers: servers}
}

// Execute returns all active servers.
func (uc *ListUseCase) Execute(ctx context.Context) ([]ServerDTO, error) {
	list, err := uc.servers.ListActive(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]ServerDTO, len(list))
	for i, s := range list {
		result[i] = ServerDTO{
			ID:     s.ID.String(),
			Name:   s.Name,
			Region: s.Region,
			Host:   s.Host,
			Port:   s.Port,
		}
	}
	return result, nil
}
