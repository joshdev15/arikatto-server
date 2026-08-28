package models

import "net/http"

type Route struct {
	Path    string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

type StaticRoute struct {
	PathPrefix  string
	StripPrefix string
	Dir         string
}

type RouteList []Route

type StaticList []StaticRoute
