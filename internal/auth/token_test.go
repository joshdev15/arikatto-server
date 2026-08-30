package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"testing"
	"time"

	"arikatto/internal/config"
)

func generateTestRSAKeys(t *testing.T) (string, string) {
	t.Helper()
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("failed to generate rsa key: %v", err)
	}

	privPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privKey),
	})

	pubBytes, err := x509.MarshalPKIXPublicKey(&privKey.PublicKey)
	if err != nil {
		t.Fatalf("failed to marshal public key: %v", err)
	}

	pubPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubBytes,
	})

	return string(privPEM), string(pubPEM)
}

func TestTokenManager_GenerateAndValidate(t *testing.T) {
	privPEM, pubPEM := generateTestRSAKeys(t)

	cfg := &config.AuthConfig{
		PrivateKeyPEM: privPEM,
		PublicKeyPEM:  pubPEM,
		Duration:      2 * time.Hour,
	}

	tm, err := NewTokenManager(cfg)
	if err != nil {
		t.Fatalf("NewTokenManager error: %v", err)
	}

	username := "joshdev"
	token, err := tm.GenerateToken(username)
	if err != nil {
		t.Fatalf("GenerateToken error: %v", err)
	}

	if token == "" {
		t.Fatal("expected non-empty token")
	}

	claim, err := tm.ValidateToken(token)
	if err != nil {
		t.Fatalf("ValidateToken error: %v", err)
	}

	if claim.User != username {
		t.Errorf("expected user %q, got %q", username, claim.User)
	}
}
