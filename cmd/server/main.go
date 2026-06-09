package main

import (
	"fmt"
	"net/http"

	"github.com/romulo/go-finance-api/internal/config"
	"github.com/romulo/go-finance-api/internal/database"
	"github.com/romulo/go-finance-api/internal/routes"
)

func main() {
	database.Connect()

	routes.RegisterRoutes()

	fmt.Printf("Servidor iniciado na porta %s\n", config.ServerPort)
	if err := http.ListenAndServe(config.ServerPort, nil); err != nil {
		fmt.Printf("Erro ao iniciar servidor: %v\n", err)
	}
}