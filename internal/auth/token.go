package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TokenManager handles JWT generation and validation using RSA keys
type TokenManager struct {
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
}

// NewTokenManager loads RSA key files and returns a TokenManager instance
func NewTokenManager(privateFile, publicFile string) (*TokenManager, error) {
	privateBytes, err := os.ReadFile(privateFile)
	if err != nil {
		return nil, fmt.Errorf("reading private key file %q: %w", privateFile, err)
	}

	publicBytes, err := os.ReadFile(publicFile)
	if err != nil {
		return nil, fmt.Errorf("reading public key file %q: %w", publicFile, err)
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		return nil, fmt.Errorf("parsing RSA private key: %w", err)
	}

	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		return nil, fmt.Errorf("parsing RSA public key: %w", err)
	}

	return &TokenManager{
		signKey:   signKey,
		verifyKey: verifyKey,
	}, nil
}

// GenerateToken creates and signs a new JWT token for the specified user
func (tm *TokenManager) GenerateToken(user string) (string, error) {
	claim := Claim{
		User:      user,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "Arikatto",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	return token.SignedString(tm.signKey)
}

// ValidateToken validates a JWT string and returns the parsed Claims
func (tm *TokenManager) ValidateToken(tokenStr string) (*Claim, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claim{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return tm.verifyKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claim, ok := token.Claims.(*Claim)
	if !ok {
		return nil, errors.New("the claims could not be parsed")
	}

	return claim, nil
}
