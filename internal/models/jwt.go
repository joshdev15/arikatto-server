package models

import "github.com/golang-jwt/jwt/v5"

// Claim structure for jwt
type Claim struct {
	User string `json:"user"`
	jwt.RegisteredClaims
}
