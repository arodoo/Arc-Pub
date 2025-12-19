// File: server_repo.go
// Purpose: Implements ServerRepository interface using PostgreSQL with sqlc.
// Provides ListActive for getting available servers and GetByID for server
// lookup. Uses pgx connection pool for database access. Converts sqlc types
// to domain entities. Used during server selection in user onboarding.
// Path: server/internal/infra/postgres/server_repo.go
// All Rights Reserved. Arc-Pub.

package postgres

import (
	"context"

	domainServer "github.com/arc-pub/server/internal/domain/server"
	"github.com/arc-pub/server/internal/infra/postgres/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ServerRepo implements ServerRepository.
type ServerRepo struct {
	queries *sqlc.Queries
}

// NewServerRepo creates a ServerRepo.
func NewServerRepo(pool *pgxpool.Pool) *ServerRepo {
	return &ServerRepo{queries: sqlc.New(pool)}
}

// ListActive returns all active servers.
func (r *ServerRepo) ListActive(ctx context.Context) ([]*domainServer.Server, error) {
	rows, err := r.queries.ListActiveServers(ctx)
	if err != nil {
		return nil, err
	}

	servers := make([]*domainServer.Server, len(rows))
	for i, row := range rows {
		servers[i] = r.rowToServer(row)
	}
	return servers, nil
}

// GetByID returns server by ID.
func (r *ServerRepo) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*domainServer.Server, error) {
	row, err := r.queries.GetServerByID(ctx, uuidToPgtype(id))
	if err != nil {
		return nil, err
	}
	return r.rowToServerByID(row), nil
}

func (r *ServerRepo) rowToServer(row sqlc.ListActiveServersRow) *domainServer.Server {
	var host *string
	var port *int
	if row.Host.Valid {
		host = &row.Host.String
	}
	if row.Port.Valid {
		p := int(row.Port.Int32)
		port = &p
	}
	return &domainServer.Server{
		ID:     pgtypeToUUID(row.ID),
		Name:   row.Name,
		Region: row.Region,
		Host:   host,
		Port:   port,
	}
}

func (r *ServerRepo) rowToServerByID(row sqlc.GetServerByIDRow) *domainServer.Server {
	var host *string
	var port *int
	if row.Host.Valid {
		host = &row.Host.String
	}
	if row.Port.Valid {
		p := int(row.Port.Int32)
		port = &p
	}
	return &domainServer.Server{
		ID:     pgtypeToUUID(row.ID),
		Name:   row.Name,
		Region: row.Region,
		Host:   host,
		Port:   port,
	}
}
