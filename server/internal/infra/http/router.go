// Arc-Pub - Metaverso 2D MMO Social
// Copyright (c) 2024. MIT License.

// Package http provides HTTP server setup.
package http

import (
	"github.com/arc-pub/server/internal/infra/http/auth"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewRouter creates a Chi router with all routes.
func NewRouter(authHandler *auth.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", authHandler.Login)
		})
	})

	return r
}
