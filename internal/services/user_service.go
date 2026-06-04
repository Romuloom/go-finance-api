package services

import (
	"github.com/romulo/go-finance-api/internal/models"
	"github.com/romulo/go-finance-api/internal/repositories"
)


func CreateUser(user models.User) models.User {
	return  repositories.SaveUser(user)
}

func ListUsers() []models.User {
	return repositories.ListUsers()
}