// File: handler.go
// Purpose: HTTP handler for authentication endpoints following Chi router
// conventions. Translates HTTP requests into application use case calls and
// formats responses. Handles JSON parsing, error translation to HTTP status
// codes, and response serialization. Keeps HTTP concerns separate from business
// logic by delegating to LoginUseCase for actual authentication processing.
// Path: server/internal/infra/http/auth/handler.go
// All Rights Reserved. Arc-Pub.

package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	appAuth "github.com/arc-pub/server/internal/application/auth"
	domainAuth "github.com/arc-pub/server/internal/domain/auth"
)

// Handler handles authentication HTTP requests.
type Handler struct {
	loginUC *appAuth.LoginUseCase
}

// NewHandler creates a Handler with use case.
func NewHandler(loginUC *appAuth.LoginUseCase) *Handler {
	return &Handler{loginUC: loginUC}
}

// Login handles POST /api/v1/auth/login.
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req appAuth.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.loginUC.Execute(r.Context(), req)
	if err != nil {
		if errors.Is(err, domainAuth.ErrInvalidCredentials) {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
