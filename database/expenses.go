package database

import "splitwise-backend/models"

func InsertInExpenseTable(expense models.Expense) {
	query := "INSERT INTO EXPENSES (description, amount,added_by,paid_by,members,isequal,issettled) VALUES ($1,$2,$3,$4,$5,$6,$7)"

	_, err := db.Exec(query, expense.Description, expense.Amount, expense.User_added, expense.User_paid, expense.Members, expense.IsEqually, expense.IsSettled)
	if err != nil {
		panic(err)
	}
}
