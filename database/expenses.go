package database

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v3/log"
	"splitwise-backend/models"
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

func SettleExpense(expense models.Expense) {

}
