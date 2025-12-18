// File: ensure_db.go
// Purpose: Provides automatic database creation functionality similar to
// Hibernate's hibernate.hbm2ddl.auto=create. Connects to the default postgres
// database to check if target database exists, creates it if missing. This
// enables zero-configuration first-run setup where developers only need to
// provide credentials and the system handles all database initialization.
// Path: server/internal/infra/postgres/migrate/ensure_db.go
// All Rights Reserved. Arc-Pub.

package migrate

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
)

// EnsureDatabase creates the database if it doesn't exist.
func EnsureDatabase(ctx context.Context, connStr string) error {
	dbName := extractDBName(connStr)
	if dbName == "" {
		return fmt.Errorf("cannot extract database name from URL")
	}

	adminURL := strings.Replace(connStr, "/"+dbName, "/postgres", 1)

	conn, err := pgx.Connect(ctx, adminURL)
	if err != nil {
		return fmt.Errorf("connect to postgres: %w", err)
	}
	defer conn.Close(ctx)

	var exists bool
	err = conn.QueryRow(ctx,
		"SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname=$1)",
		dbName,
	).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		_, err = conn.Exec(ctx, fmt.Sprintf(
			"CREATE DATABASE %s", pgx.Identifier{dbName}.Sanitize(),
		))
		if err != nil {
			return fmt.Errorf("create database: %w", err)
		}
		fmt.Printf("Created database: %s\n", dbName)
	}

	return nil
}

func extractDBName(connStr string) string {
	parts := strings.Split(connStr, "/")
	if len(parts) == 0 {
		return ""
	}
	last := parts[len(parts)-1]
	if idx := strings.Index(last, "?"); idx != -1 {
		return last[:idx]
	}
	return last
}
