package database

import (
	"encoding/json"
	"fmt"
	"splitwise-backend/models"
	"strconv"
	"strings"
)

func SelectFromUsers(uid []string) []models.User {
	var query string
	if len(uid) == 0 {
		query = "SELECT * FROM users"
	} else {
		query = "SELECT * from users where uid in (" + strings.Join(uid, ",") + ")"
	}
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {

		user := models.User{}
		var owesByte []byte // Declare a pointer to a map
		var owedByte []byte // Declare a pointer to a map

		if err := rows.Scan(&user.Uid, &user.Name, &user.Balance, &owesByte, &owedByte, &user.Number); err != nil {
			panic(err)
		}
		var owesMap map[int]float64
		var owedMap map[int]float64

		owesMap, _ = parseJsonbObject(owesByte)
		owedMap, _ = parseJsonbObject(owedByte)

		user.Owes = owesMap
		user.Owed = owedMap
		users = append(users, user)
		fmt.Printf("Name: %s, Number: %s\n", user.Name, user.Number)
	}
	return users
}

type jsonbObject struct {
	Data map[string]float64 `json:"data"`
}

func parseJsonbObject(data []byte) (map[int]float64, error) {
	var result jsonbObject
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	intFloatMap := make(map[int]float64)
	for key, val := range result.Data {
		intKey, err := strconv.Atoi(key)

		if err != nil {
			continue
		}
		intFloatMap[intKey] = val
	}
	return intFloatMap, nil
}
func InsertInUserTable(user models.User) error {
	query := "INSERT INTO USERS (name, number) VALUES ($1,$2) "
	_, err := db.Query(query, user.Name, user.Number)
	if err != nil {
		return err
	}

	// defer rows.Close()
	return nil
}

func UpdateUser(user models.User) error {
	query, queryParams := buildUpdateQuery(user, "users", "uid")
	_, err := db.Exec(query, queryParams...)
	if err != nil {
		return err
	}
	return nil
}
