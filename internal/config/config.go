package config

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/rs/cors"
)

// Config holds all configuration parameters for the application
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Auth     AuthConfig
}

// ServerConfig contains HTTP server settings
type ServerConfig struct {
	Port string
	Cors bool
}

// DatabaseConfig contains PostgreSQL connection settings
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

// DSN returns the PostgreSQL connection string
func (db *DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		db.Host, db.Port, db.User, db.Password, db.Name, db.SSLMode,
	)
}

// AuthConfig contains JWT and RSA key settings
type AuthConfig struct {
	PrivateKeyPEM  string
	PublicKeyPEM   string
	PrivateKeyPath string
	PublicKeyPath  string
	Duration       time.Duration
}

// Load reads all configuration from environment variables with safe defaults
func Load() *Config {
	var cfg Config

	// Server
	cfg.Server.Port = fmt.Sprintf(":%s", getEnv("PORT", "8080"))
	cfg.Server.Cors = getEnvBool("CORS_ENABLED", true)

	// Database
	cfg.Database.Host = getEnv("DB_HOST", "localhost")
	cfg.Database.Port = getEnv("DB_PORT", "5432")
	cfg.Database.User = getEnv("DB_USER", "admin")
	cfg.Database.Password = getEnv("DB_PASSWORD", "admin")
	cfg.Database.Name = getEnv("DB_NAME", "arikatto")
	cfg.Database.SSLMode = getEnv("DB_SSLMODE", "disable")

	// Auth
	cfg.Auth.PrivateKeyPEM = os.Getenv("JWT_PRIVATE_KEY")
	cfg.Auth.PublicKeyPEM = os.Getenv("JWT_PUBLIC_KEY")
	cfg.Auth.PrivateKeyPath = getEnv("JWT_PRIVATE_KEY_PATH", "")
	cfg.Auth.PublicKeyPath = getEnv("JWT_PUBLIC_KEY_PATH", "")
	
	durationHours := getEnvInt("JWT_DURATION_HOURS", 2)
	cfg.Auth.Duration = time.Duration(durationHours) * time.Hour

	return &cfg
}

// RunServer starts the HTTP server with configured CORS and port
func RunServer(handler http.Handler, cfg *Config) {
	fmt.Printf("Starting server on port %s\n", cfg.Server.Port)

	var serverHandler http.Handler = handler
	if cfg.Server.Cors {
		serverHandler = cors.Default().Handler(handler)
	}

	log.Fatal(http.ListenAndServe(cfg.Server.Port, serverHandler))
}

func getEnv(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	boolVal, err := strconv.ParseBool(val)
	if err != nil {
		return defaultValue
	}
	return boolVal
}

func getEnvInt(key string, defaultValue int) int {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	intVal, err := strconv.Atoi(val)
	if err != nil {
		return defaultValue
	}
	return intVal
}
