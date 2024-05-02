package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v3/log"
	"reflect"
	"splitwise-backend/models"
	"strings"
)

func InsertInExpenseTable(expense models.Expense) {

	membersJson, err := json.Marshal(expense.Members)
	if err != nil {
		panic(err)
	}
	query := "INSERT INTO EXPENSES (description, amount,added_by,paid_by,members,isequal,issettled) VALUES ($1,$2,$3,$4,$5,$6,$7)"

	_, err = db.Exec(query, expense.Description, expense.Amount, expense.User_added, expense.User_paid, membersJson, expense.IsEqually, expense.IsSettled)
	if err != nil {
		panic(err)
	}
}

func UpdateInExpenseTable(expense models.Expense) error {

	query, queryParams := buildUpdateQuery(expense)
	_, err := db.Exec(query, queryParams...)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFromExpenseTable(eid int) error {
	exists, _ := CheckIdExists("expenses", "eid", eid)
	if exists {
		query := "DELETE FROM expenses WHERE eid = $1"

		_, err := db.Query(query, eid)
		if err != nil {
			panic(err)
		}
		log.Warn("Deleted Expense", eid)
	} else {
		return errors.New("Expense not found")
	}
	return nil
}

func buildUpdateQuery(model interface{}) (string, []interface{}) {
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

	query := fmt.Sprintf("UPDATE expenses SET %s WHERE eid = $%d", strings.Join(queryVars, ", "), len(queryParams)+1)
	queryParams = append(queryParams, v.FieldByName("Eid").Interface())

	return query, queryParams
}
