package api

import (
	mw "arikatto/internal/middleware"
	"arikatto/internal/models"
	"errors"
	"net/http"
)

var (
	Welcome = models.RouteList{
		models.Route{
			Path:    "",
			Method:  http.MethodGet,
			Handler: mw.Log(greet),
		},
		models.Route{
			Path:    "/error",
			Method:  http.MethodGet,
			Handler: mw.Log(intencionalMistake),
		},
	}
)

func greet(w http.ResponseWriter, _ *http.Request) {
	resp := models.NewResponse(nil, "Arikatto", nil, http.StatusOK)
	resp.Send(w)
}

func intencionalMistake(w http.ResponseWriter, _ *http.Request) {
	err := errors.New("error")
	resp := models.NewResponse(nil, "", err, http.StatusInternalServerError)
	resp.Send(w)
}
