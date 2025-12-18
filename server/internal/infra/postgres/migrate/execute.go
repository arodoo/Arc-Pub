// File: execute.go
// Purpose: Contains the migration execution logic separated from runner.go
// for line count compliance. Handles individual migration file execution
// within a database transaction for atomicity. Checks if migration was
// already applied, reads SQL content, executes within transaction, records
// applied version in schema_migrations table, and logs success.
// Path: server/internal/infra/postgres/migrate/execute.go
// All Rights Reserved. Arc-Pub.

package migrate

import (
	"context"
	"log"
	"strings"
)

func (r *Runner) runMigration(ctx context.Context, file string) error {
	version := strings.TrimSuffix(file, ".up.sql")

	var exists bool
	err := r.pool.QueryRow(ctx,
		"SELECT EXISTS(SELECT 1 FROM schema_migrations WHERE version=$1)",
		version,
	).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	content, err := migrations.ReadFile(file)
	if err != nil {
		return err
	}

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if _, err := tx.Exec(ctx, string(content)); err != nil {
		return err
	}

	if _, err := tx.Exec(ctx,
		"INSERT INTO schema_migrations (version) VALUES ($1)",
		version,
	); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	log.Printf("Applied migration: %s", version)
	return nil
}
