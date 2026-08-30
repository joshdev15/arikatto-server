package main

import (
	"arikatto/internal/api"
	"arikatto/internal/auth"
	"arikatto/internal/config"

	_ "github.com/lib/pq"
)

func main() {
	config.LoadConfiguration()
	auth.LoadCertificates()
	server := api.Server()
	config.RunServer(server)
}
