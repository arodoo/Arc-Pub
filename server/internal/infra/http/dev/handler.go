// File: handler.go
// Purpose: Development-only HTTP handlers for testing. ResetUser clears user
// progress (faction, server, ships) to allow re-testing the onboarding flow.
// These endpoints should be disabled in production via config flag.
// Path: server/internal/infra/http/dev/handler.go
// All Rights Reserved. Arc-Pub.

package dev

import (
	"context"
	"errors"
	"net/http"

	"github.com/arc-pub/server/internal/infra/postgres/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Handler handles dev HTTP requests.
type Handler struct {
	queries *sqlc.Queries
}

// NewHandler creates a dev Handler.
func NewHandler(pool *pgxpool.Pool) *Handler {
	return &Handler{queries: sqlc.New(pool)}
}

// ResetUser handles POST /api/v1/dev/reset.
func (h *Handler) ResetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserIDFromContext(r)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	ctx := r.Context()
	if err := h.resetProgress(ctx, userID); err != nil {
		http.Error(w, "reset failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"reset"}`))
}

func (h *Handler) resetProgress(ctx context.Context, userID uuid.UUID) error {
	pgID := pgtype.UUID{Bytes: userID, Valid: true}
	if err := h.queries.DeleteUserShips(ctx, pgID); err != nil {
		return err
	}
	return h.queries.ResetUserProgress(ctx, pgID)
}

func getUserIDFromContext(r *http.Request) (uuid.UUID, error) {
	idStr := r.Header.Get("X-User-ID")
	if idStr == "" {
		return uuid.Nil, errors.New("no user id")
	}
	return uuid.Parse(idStr)
}
