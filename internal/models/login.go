package models

// Login structure for the body object
// of the login request
type Login struct {
	User     string `json:"user"`
	Password string `json:"password"`
}
