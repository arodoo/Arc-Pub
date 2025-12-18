// File: router.go
// Purpose: Configures the Chi HTTP router with middleware stack and route
// definitions. Sets up logging, panic recovery, and request ID middleware for
// observability. Defines API versioning structure (/api/v1) and mounts auth
// and user routes. Centralizes all HTTP routing configuration in one place.
// Path: server/internal/infra/http/router.go
// All Rights Reserved. Arc-Pub.

package http

import (
	"github.com/arc-pub/server/internal/infra/http/auth"
	"github.com/arc-pub/server/internal/infra/http/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewRouter creates a Chi router with all routes.
func NewRouter(authHandler *auth.Handler, userHandler *user.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", authHandler.Login)
		})
		r.Route("/user", func(r chi.Router) {
			r.Get("/profile", userHandler.GetProfile)
			r.Post("/faction", userHandler.SelectFaction)
		})
	})

	return r
}
