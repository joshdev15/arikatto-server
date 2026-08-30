package api

import (
	"arikatto/internal/auth"
	mw "arikatto/internal/middleware"
	"arikatto/internal/models"
	"encoding/json"
	"errors"
	"net/http"
)

var (
	loginPath   = "/login"
	loginRoutes = models.RouteList{
		models.Route{
			Path:    loginPath,
			Method:  http.MethodPost,
			Handler: mw.Log(login),
		},
	}
)

// login function
// this function is in charge of receiving, analyzing and generating
// a jwt token for access to the system
func login(w http.ResponseWriter, r *http.Request) {
	var data *models.Login

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		err = errors.New("forbidden method")
		resp := models.NewResponse(nil, "", err, http.StatusInternalServerError)
		resp.Send(w)
		return
	}

	//if !data.IsLoginValid() {
	//	err = errors.New("invalid data")
	//	resp := models.NewResponse(nil, "", err, http.StatusInternalServerError)
	//	resp.Send(w)
	//	return
	//}

	token, err := auth.GenerateToken(data)
	if err != nil {
		resp := models.NewResponse(nil, "", err, http.StatusInternalServerError)
		resp.Send(w)
		return
	}

	dataToken := map[string]string{"Token": token}
	resp := models.NewResponse(dataToken, "", err, http.StatusOK)
	resp.Send(w)
}
