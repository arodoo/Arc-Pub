// Arc-Pub - Metaverso 2D MMO Social
// Copyright (c) 2024. MIT License.

// Package main is the API server entry point.
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/arc-pub/server/internal/application/auth"
	"github.com/arc-pub/server/internal/config"
	"github.com/arc-pub/server/internal/infra/crypto"
	httpPkg "github.com/arc-pub/server/internal/infra/http"
	authHandler "github.com/arc-pub/server/internal/infra/http/auth"
	"github.com/arc-pub/server/internal/infra/postgres"
	"github.com/arc-pub/server/internal/infra/postgres/migrate"
	"github.com/arc-pub/server/internal/infra/token"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Auto-create database if not exists
	if err := migrate.EnsureDatabase(ctx, cfg.DatabaseURL); err != nil {
		log.Printf("warn: ensure database: %v", err)
	}

	pool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer pool.Close()

	// Auto-run migrations on startup
	runner := migrate.NewRunner(pool)
	if err := runner.Run(ctx); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	hasher := crypto.NewBcryptHasher()
	userRepo := postgres.NewUserRepo(pool)
	jwtSvc := token.NewJWTService(cfg.JWTSecret)

	seeder := postgres.NewSeeder(userRepo, hasher)
	if err := seeder.SeedAdmin(ctx); err != nil {
		log.Fatalf("failed to seed admin: %v", err)
	}

	loginUC := auth.NewLoginUseCase(userRepo, jwtSvc, hasher)
	handler := authHandler.NewHandler(loginUC)

	router := httpPkg.NewRouter(handler)

	addr := ":" + cfg.Port
	log.Printf("Server starting on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
