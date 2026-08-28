package main

import (
	"arikatto/internal/api"
	"arikatto/internal/config"

	_ "github.com/lib/pq"
)

func main() {
	config.LoadConfiguration()
	server := api.Server()
	config.RunServer(server)
}
