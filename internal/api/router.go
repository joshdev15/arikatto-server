package api

import (
	"fmt"
	"net/http"
)

const apiPrefix = "/api"

// Router wraps http.ServeMux and handles modular route registration
type Router struct {
	mux *http.ServeMux
}

// NewRouter initializes and returns a new Router
func NewRouter() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

// RegisterRoutes registers a slice of routes into the router's ServeMux
func (r *Router) RegisterRoutes(routes RouteList) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s%s", route.Method, apiPrefix, route.Path)
		r.mux.HandleFunc(pattern, route.Handler)
	}
}

// RegisterModule registers all routes from a Routable module
func (r *Router) RegisterModule(module Routable) {
	r.RegisterRoutes(module.Routes())
}

// Handler returns the underlying http.Handler
func (r *Router) Handler() http.Handler {
	return r.mux
}

// Mux returns the underlying *http.ServeMux
func (r *Router) Mux() *http.ServeMux {
	return r.mux
}
