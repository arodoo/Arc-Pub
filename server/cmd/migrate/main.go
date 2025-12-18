// Arc-Pub - Metaverso 2D MMO Social
// Copyright (c) 2024. MIT License.

// cmd/migrate: Runs database migrations.
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

	// Auto-create database if not exists
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
