package api

import "net/http"

// Route defines an HTTP endpoint structure
type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

// RouteList represents a collection of routes
type RouteList []Route

// Routable is an interface implemented by modules that expose routes
type Routable interface {
	Routes() RouteList
}
