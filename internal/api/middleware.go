package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const AuthorizationHeader = "Authorization"

// ValidatorFunc is a function type for validating a token string
type ValidatorFunc func(token string) error

// Log logs the incoming request path and method to the console
func Log(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request: %q, Method: %q\n", r.URL.Path, r.Method)
		next(w, r)
	}
}

// Authenticate returns a middleware that validates JWT tokens
func Authenticate(validate ValidatorFunc) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get(AuthorizationHeader)
			if token == "" {
				Error(w, http.StatusUnauthorized, errors.New("missing authorization header"))
				return
			}

			// Support "Bearer <token>" or raw token
			token = strings.TrimPrefix(token, "Bearer ")

			if err := validate(token); err != nil {
				Error(w, http.StatusUnauthorized, errors.New("not authorized"))
				return
			}

			next(w, r)
		}
	}
}
