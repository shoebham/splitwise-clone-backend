package handlers

import (
	"splitwise-backend/database"
	"splitwise-backend/models"
)

func GetAllExpense() {

	database.GetAllData("expenses")

}

func CreateExpense(expense models.Expense) error {
	if err := database.InsertInExpenseTable(expense); err != nil {
		return err
	}
	return nil
}

func UpdateExpense(expense models.Expense) error {
	if err := database.UpdateInExpenseTable(expense); err != nil {
		return err
	}
	return nil
}

func DeleteExpense(eid int) error {

	if err := database.DeleteFromExpenseTable(eid); err != nil {
		return err
	}
	return nil
}

func SettleExpense(expense models.Expense) error {

	if err := database.UpdateInExpenseTable(expense); err != nil {
		return err
	}
	return nil
}
