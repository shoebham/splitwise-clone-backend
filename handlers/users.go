package handlers

import (
	"splitwise-backend/database"
	"splitwise-backend/models"
)

func CreateUser(user models.User) error {
	err := database.InsertInUserTable(user)
	if err != nil {
		return err
	}
	return nil
}

func GetAllUsers() []models.User {
	return database.SelectFromUsers(nil)
}
func GetUserById(id []string) []models.User {
	return database.SelectFromUsers(id)
}
