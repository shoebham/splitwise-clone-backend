package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v3/log"
	"splitwise-backend/models"
	"strings"
)

func InsertInExpenseTable(expense models.Expense) error {

	membersJson, err := json.Marshal(expense.Members)
	if err != nil {
		return err
	}
	query := "INSERT INTO EXPENSES (description, amount,added_by,paid_by,members,isequal,issettled) VALUES ($1,$2,$3,$4,$5,$6,$7)"

	_, err = db.Exec(query, expense.Description, expense.Amount, expense.UserAdded, expense.UserPaid, membersJson, expense.IsEqually, expense.IsSettled)
	if err != nil {
		return err
	}
	return nil
}

func UpdateInExpenseTable(expense models.Expense) error {

	query, queryParams := buildUpdateQuery(expense, "expenses", "eid")
	// Replace placeholders in the query string with actual values
	var replacedQuery string
	replacedQuery = query
	for i, param := range queryParams {
		replacedQuery = strings.Replace(replacedQuery, fmt.Sprintf("$%d", i+1), fmt.Sprintf("%v", param), -1)
	}

	fmt.Println("Replaced query:", replacedQuery)

	_, err := db.Exec(replacedQuery)
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

func SettleExpense(expense models.Expense) {

}
