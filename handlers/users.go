package handlers

import (
	"splitwise-backend/database"
	"splitwise-backend/models"
)

func CreateUser(user models.User) {
	database.InsertInUserTable(user)
}

func GetAllUsers() []models.User {
	return database.SelectFromUsers(nil)
}
func GetUserById(id []string) []models.User {
	return database.SelectFromUsers(id)
}
