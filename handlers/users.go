package handlers

import (
	"fmt"
	"splitwise-backend/database"
	"splitwise-backend/models"
)

func CreateUser(user models.User) error {
	preInsertChecks(&user)
	err := database.InsertInUserTable(user)
	if err != nil {
		return err
	}
	return nil
}
func UpdateUser(user models.User) error {
	preInsertChecks(&user)
	if err := database.UpdateUser(user); err != nil {
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

func preInsertChecks(user *models.User) {

	bal := 0.0
	fmt.Printf("User: %v, Balance: %v Owes: %v \n", user.Name, bal, user.Owes)
	for _, share := range user.Owes {
		bal -= share
	}

	for _, share := range user.Owed {
		bal += share
	}
	user.Balance = bal

}
