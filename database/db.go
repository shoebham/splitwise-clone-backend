package database

import (
	"database/sql"
	"fmt"
	"reflect"
	"splitwise-backend/constants"
	"strings"

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

func buildUpdateQuery(model interface{}, tableName string, id string) (string, []interface{}) {
	var queryVars []string
	var queryParams []interface{}
	v := reflect.ValueOf(model)
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		fieldValue := v.Field(i).Interface()

		// Check if fieldValue is a map[int]float64
		if m, ok := fieldValue.(map[int]float64); ok {
			// Handle map[int]float64 separately
			for key, value := range m {
				queryVars = append(queryVars, fmt.Sprintf("%s[%d] = $%d", field.Tag.Get("json"), key, len(queryParams)+1))
				queryParams = append(queryParams, value)
			}
			continue // Skip the rest of the loop iteration
		}

		if fieldValue != reflect.Zero(v.Field(i).Type()).Interface() {
			queryVars = append(queryVars, fmt.Sprintf("%s = $%d", field.Tag.Get("json"), len(queryParams)+1))
			queryParams = append(queryParams, fieldValue)
		}

	}

	idUpprCase := strings.Title(id)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s = $%d", tableName, strings.Join(queryVars, ", "), id, len(queryParams)+1)
	queryParams = append(queryParams, v.FieldByName(idUpprCase).Interface())

	return query, queryParams
}
