package api

import (
	"arikatto/internal/models"
	"fmt"
	"net/http"
)

var (
	r       = http.NewServeMux()
	apiPath = "/api"
)

// Server function
// This function adds the routes by category
// api and static then returns the router object
// with all the configured routes.
// Returns:
// - r *mux.Router
func Server() *http.ServeMux {
	addRoutes(Welcome)
	return r
}

// addRoutes function
// This function iterates over a list of routes
// and adds them to the router object.
// Params:
// - routesList model.List
func addRoutes(routesList models.RouteList) {
	for _, v := range routesList {
		path := fmt.Sprintf("%s %s%s", v.Method, apiPath, v.Path)
		r.HandleFunc(path, v.Handler)
	}
}

// addStaticRoutes function
// This function iterates over a list of static routes
// and adds them to the router object.
// Params:
// - routesList model.StaticList
//func addStaticRoutes(mux *http.ServeMux, route models.StaticRoute) {
//	fileServer := http.FileServer(http.Dir(route.Dir))
//	handler := http.StripPrefix(route.StripPrefix, fileServer)
//	pattern := fmt.Sprintf("GET %s", route.PathPrefix)
//
//	if !strings.HasSuffix(pattern, "/") {
//		pattern += "/"
//	}
//
//	mux.Handle(pattern, handler)
//}
