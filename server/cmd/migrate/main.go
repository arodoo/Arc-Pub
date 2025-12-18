// File: main.go
// Purpose: Standalone CLI tool for running database migrations independently
// from the main server. Creates the target database if it does not exist,
// then applies all pending migrations in order. Useful for CI/CD pipelines
// and manual database setup without starting the full application server.
// Path: server/cmd/migrate/main.go
// All Rights Reserved. Arc-Pub.

package main

import (
	"context"
	"log"

	"github.com/arc-pub/server/internal/config"
	"github.com/arc-pub/server/internal/infra/postgres/migrate"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	if err := migrate.EnsureDatabase(ctx, cfg.DatabaseURL); err != nil {
		log.Fatalf("ensure database: %v", err)
	}

	pool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("connect: %v", err)
	}
	defer pool.Close()

	runner := migrate.NewRunner(pool)
	if err := runner.Run(ctx); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	log.Println("Migrations completed successfully")
}
