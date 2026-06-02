package routes

import (
	"net/http"

	"github.com/romulo/go-finance-api/internal/handlers"
)

func RegisterRoutes(){
	http.HandleFunc("/health", handlers.HealthHandler)
}