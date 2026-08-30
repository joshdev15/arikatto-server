package config

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
	"gopkg.in/yaml.v3"
)

const defaultConfigFile = "config/server.yml"

// Config represents the application configuration
type Config struct {
	Port   string
	Server struct {
		Cors             bool `yaml:"cors"`
		CertificatesPath struct {
			Private string `yaml:"private"`
			Public  string `yaml:"public"`
		} `yaml:"certificates"`
	} `yaml:"server"`
}

// Load loads the configuration from YAML file and environment variables
func Load() (*Config, error) {
	var cfg Config

	data, err := os.ReadFile(defaultConfigFile)
	if err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("unmarshaling config yaml: %w", err)
	}

	envPort := os.Getenv("PORT")
	if envPort == "" {
		envPort = "8080"
	}
	cfg.Port = fmt.Sprintf(":%s", envPort)

	return &cfg, nil
}

// RunServer starts the HTTP server using the provided handler and configuration
func RunServer(handler http.Handler, cfg *Config) {
	fmt.Printf("Starting server on port %s\n", cfg.Port)

	var serverHandler http.Handler = handler
	if cfg.Server.Cors {
		serverHandler = cors.Default().Handler(handler)
	}

	log.Fatal(http.ListenAndServe(cfg.Port, serverHandler))
}
