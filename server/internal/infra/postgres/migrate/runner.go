// File: runner.go
// Purpose: Provides database migration functionality with version tracking
// similar to Flyway or Liquibase. Embeds SQL migration files at compile time
// using Go embed directive. Tracks applied migrations in schema_migrations
// table to prevent re-execution. Runs migrations in sorted order within
// transactions for atomicity. Supports multiple migration files following
// the naming convention XXX_description.up.sql.
// Path: server/internal/infra/postgres/migrate/runner.go
// All Rights Reserved. Arc-Pub.

package migrate

import (
	"context"
	"embed"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed *.sql
var migrations embed.FS

// Runner handles database migrations.
type Runner struct {
	pool *pgxpool.Pool
}

// NewRunner creates a migration runner.
func NewRunner(pool *pgxpool.Pool) *Runner {
	return &Runner{pool: pool}
}

// Run executes all pending migrations.
func (r *Runner) Run(ctx context.Context) error {
	if err := r.createMigrationsTable(ctx); err != nil {
		return fmt.Errorf("create migrations table: %w", err)
	}

	files, err := r.getMigrationFiles()
	if err != nil {
		return fmt.Errorf("get migration files: %w", err)
	}

	for _, file := range files {
		if err := r.runMigration(ctx, file); err != nil {
			return fmt.Errorf("run %s: %w", file, err)
		}
	}

	return nil
}

func (r *Runner) createMigrationsTable(ctx context.Context) error {
	query := `
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version VARCHAR(255) PRIMARY KEY,
			applied_at TIMESTAMPTZ DEFAULT NOW()
		)`
	_, err := r.pool.Exec(ctx, query)
	return err
}

func (r *Runner) getMigrationFiles() ([]string, error) {
	entries, err := migrations.ReadDir(".")
	if err != nil {
		return nil, err
	}

	var files []string
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".up.sql") {
			files = append(files, e.Name())
		}
	}
	sort.Strings(files)
	return files, nil
}

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
