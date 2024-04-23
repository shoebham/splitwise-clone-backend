package database

import (
	"fmt"
	"splitwise-backend/models"
)

func SelectFromUsers() []models.User {
	query := "SELECT * from users"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {

		user := models.User{}
		var owesMap *map[*models.User]float64 // Declare a pointer to a map
		var owedMap *map[*models.User]float64 // Declare a pointer to a map

		if err := rows.Scan(&user.Uid, &user.Name, &user.Balance, &owesMap, &owedMap, &user.Number); err != nil {
			panic(err)
		}
		users = append(users, user)
		fmt.Printf("Name: %s, Number: %s\n", user.Name, user.Number)
	}
	return users
}
func InsertInUserTable(user models.User) {
	query := "INSERT INTO USERS (name, number) VALUES ($1,$2) "
	_, err := db.Query(query, user.Name, user.Number)
	if err != nil {
		panic(err)
	}

	// defer rows.Close()
}
