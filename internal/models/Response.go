package models

import (
	"encoding/json"
	"net/http"
)

// Response structure of a request
type Response struct {
	MsgType    string      `json:"msgType,omitempty"`
	Msg        string      `json:"msg,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Err        string      `json:"error,omitempty"`
	StatusCode int         `json:"statusCode,omitempty"`
}

// NewResponse Public function to create a Respose instance more easily
func NewResponse(data interface{}, msg string, err error, statusCode int) *Response {
	errorMessage := ""
	if err != nil {
		errorMessage = err.Error()
	}

	return &Response{
		MsgType:    http.StatusText(statusCode),
		Msg:        msg,
		Data:       data,
		Err:        errorMessage,
		StatusCode: statusCode,
	}
}

// Send Public method for submitting a petition response
func (res *Response) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.StatusCode)

	//colorlog.Action(res.StatusCode)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
