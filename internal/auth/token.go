package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"os"
	"time"

	"arikatto/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

// TokenManager handles JWT generation and validation using RSA keys
type TokenManager struct {
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
	duration  time.Duration
}

// NewTokenManager creates a TokenManager from AuthConfig
func NewTokenManager(cfg *config.AuthConfig) (*TokenManager, error) {
	privateBytes, err := loadKeyBytes(cfg.PrivateKeyPEM, cfg.PrivateKeyPath, "private")
	if err != nil {
		return nil, err
	}

	publicBytes, err := loadKeyBytes(cfg.PublicKeyPEM, cfg.PublicKeyPath, "public")
	if err != nil {
		return nil, err
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		return nil, fmt.Errorf("parsing RSA private key: %w", err)
	}

	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		return nil, fmt.Errorf("parsing RSA public key: %w", err)
	}

	duration := cfg.Duration
	if duration <= 0 {
		duration = 2 * time.Hour
	}

	return &TokenManager{
		signKey:   signKey,
		verifyKey: verifyKey,
		duration:  duration,
	}, nil
}

func loadKeyBytes(pemContent, filePath, keyType string) ([]byte, error) {
	if pemContent != "" {
		return []byte(pemContent), nil
	}

	if filePath != "" {
		bytes, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("reading %s key file %q: %w", keyType, filePath, err)
		}
		return bytes, nil
	}

	return nil, fmt.Errorf("no %s key provided (set JWT_%s_KEY or JWT_%s_KEY_PATH)", keyType, keyType, keyType)
}

// GenerateToken creates and signs a new JWT token for the specified user
func (tm *TokenManager) GenerateToken(user string) (string, error) {
	claim := Claim{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tm.duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "Arikatto",
		},
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
