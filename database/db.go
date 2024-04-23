package database

import (
	"database/sql"
	"fmt"
	"splitwise-backend/constants"

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

func CheckIdExists(tableName string, columnName string, id int) (bool, error) {

	query := fmt.Sprintf("SELECT exists(select 1 from %s where %s = %d)", tableName, columnName, id)
	var exists bool
	err := db.QueryRow(query).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil

}
