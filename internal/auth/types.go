package auth

import "github.com/golang-jwt/jwt/v5"

// Login represents credentials for authentication
type Login struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

// Claim represents the JWT claims payload
type Claim struct {
	User string `json:"user"`
	jwt.RegisteredClaims
}
