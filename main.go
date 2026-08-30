package main

import (
	"log"

	"arikatto/internal/api"
	"arikatto/internal/auth"
	"arikatto/internal/config"

	_ "github.com/lib/pq"
)

func main() {
	// 1. Cargar configuración
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// 2. Inicializar gestor de tokens / certificados RSA
	tokenManager, err := auth.NewTokenManager(
		cfg.Server.CertificatesPath.Private,
		cfg.Server.CertificatesPath.Public,
	)
	if err != nil {
		log.Fatalf("Error loading certificates: %v", err)
	}

	// 3. Inicializar módulos y handlers
	welcomeHandler := api.NewWelcomeHandler()
	authHandler := auth.NewHandler(tokenManager)

	// 4. Crear router y registrar módulos
	router := api.NewRouter()
	router.RegisterModule(welcomeHandler)
	router.RegisterModule(authHandler)

	// 5. Iniciar servidor
	config.RunServer(router.Handler(), cfg)
}
