package database

import (
	"database/sql"
	"fmt"
	"splitwise-backend/constants"
	"splitwise-backend/models"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDb() {
	connStr := fmt.Sprintf("user=%s dbname=splitwise_backend password=%s sslmode=disable", constants.DB_USER, constants.DB_PASS)
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connection succssfull")
	// defer db.Close()

}
func GetAllData(tableName string) {

	query := fmt.Sprintf("SELECT * from %s", tableName)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		fmt.Printf("%v", rows.Scan())
	}
}

func SelectFromUsers() {
	query := "SELECT * from users"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {

		user := models.User{}
		var owesMap *map[*models.User]float64 // Declare a pointer to a map
		var owedMap *map[*models.User]float64 // Declare a pointer to a map

		if err := rows.Scan(&user.UserID, &user.Name, &user.Balance, &owesMap, &owedMap, &user.Number); err != nil {
			panic(err)
		}
		fmt.Printf("Name: %s, Number: %s\n", user.Name, user.Number)
	}
}
func InsertInUserTable(user models.User) {
	query := "INSERT INTO USERS (name, number) VALUES ($1,$2)"
	_, err := db.Query(query, user.Name, user.Number)
	if err != nil {
		panic(err)
	}
	// defer rows.Close()
}

func InsertInGroupTable(group models.Group) {
	query := "INSERT INTO GROUPS (group_name, members) VALUES ($1,$2)"
	_, err := db.Exec(query, group.GroupName, pq.Array(group.Members))
	if err != nil {
		panic(err)
	}
}

func InsertInExpenseTable(expense models.Expense) {
	query := "INSERT INTO EXPENSES (description, amount,added_by,paid_by,members,isequal,issettled) VALUES ($1,$2,$3,$4,$5,$6,$7)"

	_, err := db.Exec(query, expense.Description, expense.Amount, expense.User_added, expense.User_paid, expense.Members, expense.IsEqually, expense.IsSettled)
	if err != nil {
		panic(err)
	}
}

func GetGroups() {
}
