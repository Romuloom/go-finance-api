package routes

import (
	"net/http"

	"github.com/romulo/go-finance-api/internal/handlers"
)

func RegisterRoutes(){
	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.ListUsers(w, r)
		case http.MethodPost:
			handlers.CreateUser(w, r)	
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}