package main

import (
	"fmt"
	"net/http"

	"github.com/romulo/go-finance-api/internal/config"
	"github.com/romulo/go-finance-api/internal/routes"
)

func main() {

	routes.RegisterRoutes()

	fmt.Printf("Servidor iniciado na porta %s\n", config.ServerPort)
	http.ListenAndServe(config.ServerPort, nil)
}