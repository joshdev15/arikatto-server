package config

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
	"gopkg.in/yaml.v3"
)

const (
	userFilePath = "config/server.yml"
)

var (
	port          string
	configuration configurationFile
)

func RunServer(r *http.ServeMux) {
	fmt.Printf("Starting server on port %s\n", port)

	if configuration.Server.Cors {
		handler := cors.Default().Handler(r)
		log.Fatal(http.ListenAndServe(port, handler))
		return
	}

	log.Fatal(http.ListenAndServe(port, r))
}

func loadPort() {
	defaultPort := "8080"
	envPort := os.Getenv("PORT")
	if envPort == "" {
		envPort = defaultPort
	}

	port = fmt.Sprintf(":%s", envPort)
}

func loadConfig(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &configuration)
	if err != nil {
		return err
	}

	return nil
}

func LoadConfiguration() {
	loadPort()

	err := loadConfig(userFilePath)
	if err != nil {
		log.Fatal(err)
	}
}

func Get() ConfigurationReader {
	return &configuration
}
