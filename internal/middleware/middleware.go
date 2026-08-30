package middleware

import (
	"arikatto/internal/auth"
	"arikatto/internal/models"
	"errors"
	"fmt"
	"net/http"
)

const (
	Authorization = "Authorization"
)

var (
	errMessage = "not authorized"
	defaultErr = errors.New(errMessage)
)

// Log Function
// Public function prints the path and the request
// method to the console.
func Log(run func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request: %q, Method: %q\n", r.URL.Path, r.Method)
		run(w, r)
	}
}

// Authentication Function
// Public function that verifies if the jwt present
// in the request has permission to access the resources.
func Authentication(run func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get(Authorization)
		_, err := auth.ValidationToken(token)
		if err != nil {
			forbidden(w, r)
			return
		}
		run(w, r)
	}
}

// forbidden function
// Private function that handles the action to be taken
// if the token is invalid.
func forbidden(w http.ResponseWriter, r *http.Request) {
	resp := models.NewResponse(nil, "", defaultErr, http.StatusInternalServerError)
	resp.Send(w)
}
