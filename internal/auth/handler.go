package auth

import (
	"arikatto/internal/api"
	"encoding/json"
	"errors"
	"net/http"
)

// Handler manages HTTP authentication endpoints
type Handler struct {
	tokenManager *TokenManager
}

// NewHandler creates a new auth Handler
func NewHandler(tm *TokenManager) *Handler {
	return &Handler{
		tokenManager: tm,
	}
}

// Routes returns the list of routes exposed by the auth module
func (h *Handler) Routes() api.RouteList {
	return api.RouteList{
		{
			Path:    "/login",
			Method:  http.MethodPost,
			Handler: api.Log(h.Login),
		},
	}
}

// Login handles authentication request and returns a JWT token
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var data Login

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		api.Error(w, http.StatusBadRequest, errors.New("invalid request body"))
		return
	}

	if data.User == "" || data.Password == "" {
		api.Error(w, http.StatusBadRequest, errors.New("user and password are required"))
		return
	}

	token, err := h.tokenManager.GenerateToken(data.User)
	if err != nil {
		api.Error(w, http.StatusInternalServerError, errors.New("could not generate token"))
		return
	}

	dataToken := map[string]string{"Token": token}
	api.JSON(w, http.StatusOK, dataToken, "authenticated successfully")
}
