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
func GetAllData(tableName string) error {
	rows, err := selectAllFromTableInternal(db, tableName)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Get the column names
	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	// Create a slice to hold the values
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	columnName := make([]interface{}, len(columns))

	for rows.Next() {
		// Populate the value pointers
		for i := range columns {
			valuePtrs[i] = &values[i]
			columnName[i] = columns[i]
		}

		// Scan the row into the value pointers
		if err := rows.Scan(valuePtrs...); err != nil {
			return err
		}

		// Print the row
		for i := range columns {
			fmt.Printf("[%s]: %v \n", columnName[i], values[i])
		}
		fmt.Println("--------")
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func selectAllFromTableInternal(db *sql.DB, tableName string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	return rows, nil
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
