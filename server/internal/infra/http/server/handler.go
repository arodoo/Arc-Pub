// File: handler.go
// Purpose: HTTP handlers for server listing and selection endpoints. ListServers
// returns available game servers. SelectServer handles one-time server choice.
// Extracts user ID from request headers (TODO: JWT middleware). Used during
// user onboarding before faction selection. Returns JSON responses.
// Path: server/internal/infra/http/server/handler.go
// All Rights Reserved. Arc-Pub.

package server

import (
	"encoding/json"
	"errors"
	"net/http"

	appServer "github.com/arc-pub/server/internal/application/server"
	"github.com/google/uuid"
)

// Handler handles server HTTP requests.
type Handler struct {
	listUC   *appServer.ListUseCase
	selectUC *appServer.SelectUseCase
}

// NewHandler creates a Handler with use cases.
func NewHandler(
	listUC *appServer.ListUseCase,
	selectUC *appServer.SelectUseCase,
) *Handler {
	return &Handler{listUC: listUC, selectUC: selectUC}
}

// ListServers handles GET /api/v1/servers.
func (h *Handler) ListServers(w http.ResponseWriter, r *http.Request) {
	servers, err := h.listUC.Execute(r.Context())
	if err != nil {
		http.Error(w, "failed to list servers", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(servers)
}

// SelectServer handles POST /api/v1/user/server.
func (h *Handler) SelectServer(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserIDFromContext(r)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var req appServer.ServerSelectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.selectUC.Execute(r.Context(), userID, req)
	if err != nil {
		if errors.Is(err, appServer.ErrServerNotFound) {
			http.Error(w, "server not found", http.StatusNotFound)
			return
		}
		if errors.Is(err, appServer.ErrServerAlreadySet) {
			http.Error(w, "server already set", http.StatusConflict)
			return
		}
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func getUserIDFromContext(r *http.Request) (uuid.UUID, error) {
	idStr := r.Header.Get("X-User-ID")
	if idStr == "" {
		return uuid.Nil, errors.New("no user id")
	}
	return uuid.Parse(idStr)
}
