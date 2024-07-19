package database

import (
	"database/sql"
	"fmt"
	"reflect"
	"splitwise-backend/constants"
	"strings"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDb() error {
	connStr := fmt.Sprintf("user=%s dbname=splitwise_backend password=%s sslmode=disable", constants.DB_USER, constants.DB_PASS)
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {

		return err
	}
	fmt.Println("Database connection succssfull")
	return DB.Ping()
	// defer DB.Close()

}
func GetAllData(tableName string) error {
	rows, err := selectAllFromTableInternal(DB, tableName)
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

func selectAllFromTableInternal(DB *sql.DB, tableName string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func CheckIdExists(tableName string, columnName string, id int) (bool, error) {

	query := fmt.Sprintf("SELECT exists(select 1 from %s where %s = %d)", tableName, columnName, id)
	var exists bool
	err := DB.QueryRow(query).Scan(&exists)
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
			jsonName := field.Tag.Get("sql")
			temp := fmt.Sprintf("%s=%s || '{ ", jsonName, jsonName)
			var tempQueryVars []string
			// Handle map[int]float64 separately
			for key, value := range m {

				tempQueryVars = append(tempQueryVars, fmt.Sprintf(`"%d":$%d`, key, len(queryParams)+1))
				queryParams = append(queryParams, value)

			}
			temp += strings.Join(tempQueryVars, ",")
			temp += "}'::jsonb"
			queryVars = append(queryVars, temp)
			continue // Skip the rest of the loop iteration
		}

		if fieldValue != reflect.Zero(v.Field(i).Type()).Interface() {
			jsonName := field.Tag.Get("sql")
			if (jsonName == "Eid" || jsonName == "Gid" || jsonName == "Uid") && (i == 0) {
				continue
			}
			switch fieldValue.(type) {
			case map[int]float64:
				m := fieldValue.(map[int]float64)
				for key, value := range m {
					queryVars = append(queryVars, fmt.Sprintf("%s->>%d = $%d", jsonName, key, len(queryParams)+1))
					queryParams = append(queryParams, value)
				}
			default:
				queryVars = append(queryVars, fmt.Sprintf("%s = $%d", jsonName, len(queryParams)+1))
				queryParams = append(queryParams, fieldValue)
			}
		}

	}

	idUpprCase := strings.Title(id)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s = $%d", tableName, strings.Join(queryVars, ", "), id, len(queryParams)+1)
	queryParams = append(queryParams, v.FieldByName(idUpprCase).Interface())

	return query, queryParams
}
