package auth

import (
	"arikatto/internal/models"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken Public function that generates the jwt for the user
func GenerateToken(data *models.Login) (string, error) {
	claim := models.Claim{
		User:      data.User,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "Arikatto",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)

	signedToken, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidationToken
// Public function that validates a jwt
func ValidationToken(t string) (models.Claim, error) {
	token, err := jwt.ParseWithClaims(t, &models.Claim{}, verifyFunction)
	if err != nil {
		fmt.Println("token validation", err)
		return models.Claim{}, err
	}

	if !token.Valid {
		return models.Claim{}, errors.New("invalid token")
	}

	claim, ok := token.Claims.(*models.Claim)
	if !ok {
		return models.Claim{}, errors.New("the claim could not be obtained")
	}

	return *claim, nil
}

// verifyFunction
// Function that returns the public key for verification
func verifyFunction(t *jwt.Token) (any, error) {
	return verifyKey, nil
}
