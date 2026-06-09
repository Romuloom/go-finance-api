package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/romulo/go-finance-api/internal/models"
	"github.com/romulo/go-finance-api/internal/services"
)



func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("JSON inválido"))

		return
	}



	createdUser, err := services.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

func ListUsers(w http.ResponseWriter, r *http.Request) {

	users := services.ListUsers()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}