package api

import (
	"errors"
	"net/http"
)

// WelcomeHandler handles general API root and health check endpoints
type WelcomeHandler struct{}

// NewWelcomeHandler creates a new WelcomeHandler
func NewWelcomeHandler() *WelcomeHandler {
	return &WelcomeHandler{}
}

// Routes returns the list of routes for the welcome endpoint
func (h *WelcomeHandler) Routes() RouteList {
	return RouteList{
		{
			Path:    "",
			Method:  http.MethodGet,
			Handler: Log(h.Greet),
		},
		{
			Path:    "/error",
			Method:  http.MethodGet,
			Handler: Log(h.IntentionalMistake),
		},
	}
}

// Greet responds with a greeting
func (h *WelcomeHandler) Greet(w http.ResponseWriter, _ *http.Request) {
	JSON(w, http.StatusOK, nil, "Arikatto")
}

// IntentionalMistake triggers an intentional error response for testing
func (h *WelcomeHandler) IntentionalMistake(w http.ResponseWriter, _ *http.Request) {
	err := errors.New("intentional error")
	Error(w, http.StatusInternalServerError, err)
}
