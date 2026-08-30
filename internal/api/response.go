package api

import (
	"encoding/json"
	"net/http"
)

// Response structure of a request
type Response struct {
	MsgType    string `json:"msgType,omitempty"`
	Msg        string `json:"msg,omitempty"`
	Data       any    `json:"data,omitempty"`
	Err        string `json:"error,omitempty"`
	StatusCode int    `json:"statusCode,omitempty"`
}

// NewResponse creates a new Response instance
func NewResponse(data any, msg string, err error, statusCode int) *Response {
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

// Send writes the response to the http.ResponseWriter
func (res *Response) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.StatusCode)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// JSON sends a JSON response with status code, data and message
func JSON(w http.ResponseWriter, statusCode int, data any, msg string) {
	NewResponse(data, msg, nil, statusCode).Send(w)
}

// Error sends a JSON error response with appropriate status code
func Error(w http.ResponseWriter, statusCode int, err error) {
	NewResponse(nil, "", err, statusCode).Send(w)
}
