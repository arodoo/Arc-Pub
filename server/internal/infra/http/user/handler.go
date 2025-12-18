// File: handler.go
// Purpose: HTTP handlers for user profile and faction endpoints. GetProfile
// returns user data for lobby display. SelectFaction handles one-time faction
// choice for new users. Extracts user ID from JWT context (TODO: implement
// middleware). Translates between HTTP and application layer with proper
// error responses. Uses Chi router conventions for RESTful API design.
// Path: server/internal/infra/http/user/handler.go
// All Rights Reserved. Arc-Pub.

package user

import (
	"encoding/json"
	"errors"
	"net/http"

	appUser "github.com/arc-pub/server/internal/application/user"
	"github.com/google/uuid"
)

// Handler handles user HTTP requests.
type Handler struct {
	profileUC *appUser.ProfileUseCase
	factionUC *appUser.FactionUseCase
}

// NewHandler creates a Handler with use cases.
func NewHandler(
	profileUC *appUser.ProfileUseCase,
	factionUC *appUser.FactionUseCase,
) *Handler {
	return &Handler{profileUC: profileUC, factionUC: factionUC}
}

// GetProfile handles GET /api/v1/user/profile.
func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserIDFromContext(r)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	resp, err := h.profileUC.Execute(r.Context(), userID)
	if err != nil {
		http.Error(w, "profile not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// SelectFaction handles POST /api/v1/user/faction.
func (h *Handler) SelectFaction(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserIDFromContext(r)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var req appUser.FactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.factionUC.Execute(r.Context(), userID, req)
	if err != nil {
		if errors.Is(err, appUser.ErrInvalidFaction) {
			http.Error(w, "invalid faction", http.StatusBadRequest)
			return
		}
		if errors.Is(err, appUser.ErrFactionAlreadySet) {
			http.Error(w, "faction already set", http.StatusConflict)
			return
		}
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// TODO: Replace with JWT middleware context extraction
func getUserIDFromContext(r *http.Request) (uuid.UUID, error) {
	idStr := r.Header.Get("X-User-ID")
	if idStr == "" {
		return uuid.Nil, errors.New("no user id")
	}
	return uuid.Parse(idStr)
}
