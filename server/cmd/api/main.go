// File: main.go
// Purpose: Entry point for the Arc-Pub API server. Initializes configuration,
// establishes PostgreSQL database connection, runs automatic migrations on
// startup, seeds admin user, configures auth/user/server use cases, starts
// HTTP server with Chi router. Wires all dependencies for the application.
// Path: server/cmd/api/main.go
// All Rights Reserved. Arc-Pub.

package main

import (
	"context"
	"log"
	"net/http"

	authApp "github.com/arc-pub/server/internal/application/auth"
	serverApp "github.com/arc-pub/server/internal/application/server"
	userApp "github.com/arc-pub/server/internal/application/user"
	"github.com/arc-pub/server/internal/config"
	"github.com/arc-pub/server/internal/infra/crypto"
	httpPkg "github.com/arc-pub/server/internal/infra/http"
	authHandler "github.com/arc-pub/server/internal/infra/http/auth"
	serverHandler "github.com/arc-pub/server/internal/infra/http/server"
	userHandler "github.com/arc-pub/server/internal/infra/http/user"
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

	if err := migrate.EnsureDatabase(ctx, cfg.DatabaseURL); err != nil {
		log.Printf("warn: ensure database: %v", err)
	}

	pool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer pool.Close()

	runner := migrate.NewRunner(pool)
	if err := runner.Run(ctx); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	// Repos
	hasher := crypto.NewBcryptHasher()
	userRepo := postgres.NewUserRepo(pool)
	profileRepo := postgres.NewProfileRepo(pool)
	shipRepo := postgres.NewShipRepo(pool)
	serverRepo := postgres.NewServerRepo(pool)
	userServerRepo := postgres.NewUserServerRepo(pool)
	jwtSvc := token.NewJWTService(cfg.JWTSecret)

	// Seeder
	seeder := postgres.NewSeeder(userRepo, hasher)
	if err := seeder.SeedAdmin(ctx); err != nil {
		log.Fatalf("failed to seed admin: %v", err)
	}

	// Use Cases
	loginUC := authApp.NewLoginUseCase(userRepo, jwtSvc, hasher)
	profileUC := userApp.NewProfileUseCase(profileRepo, shipRepo)
	factionUC := userApp.NewFactionUseCase(profileRepo, shipRepo)
	listServersUC := serverApp.NewListUseCase(serverRepo)
	selectServerUC := serverApp.NewSelectUseCase(serverRepo, userServerRepo)

	// Handlers
	authH := authHandler.NewHandler(loginUC)
	userH := userHandler.NewHandler(profileUC, factionUC)
	serverH := serverHandler.NewHandler(listServersUC, selectServerUC)

	router := httpPkg.NewRouter(authH, userH, serverH)

	addr := ":" + cfg.Port
	log.Printf("Server starting on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
