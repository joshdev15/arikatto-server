package main

import (
	"log"

	"arikatto/internal/api"
	"arikatto/internal/auth"
	"arikatto/internal/config"

	_ "github.com/lib/pq"
)

func main() {
	// 1. Cargar configuración desde variables de entorno
	cfg := config.Load()

	// 2. Inicializar gestor de tokens / certificados RSA
	tokenManager, err := auth.NewTokenManager(&cfg.Auth)
	if err != nil {
		log.Fatalf("Error initializing auth tokens: %v", err)
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
