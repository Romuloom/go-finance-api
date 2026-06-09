package services

import (
	"errors"
	"strings"

	"github.com/romulo/go-finance-api/internal/models"
	"github.com/romulo/go-finance-api/internal/repositories"
)

func CreateUser(user models.User) (models.User, error) {

	if strings.TrimSpace(user.Name) == "" {
		return models.User{}, errors.New("nome é obrigatório")
	}
	if strings.TrimSpace(user.Email) == "" {
		return models.User{}, errors.New("email é obrigatório")
	}

	return repositories.SaveUser(user), nil
}

func ListUsers() []models.User {
	return repositories.ListUsers()
}
