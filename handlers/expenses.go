package handlers

import (
	"fmt"
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

	// for each user update in user table
	// if not equal then iterate over user members, check the sum of share it should be equal to amount
	// for each user update users share, if user paid and in members, do nothing for that user

	err := updateRelatedTables(expense)
	if err != nil {
		return err
	}
	fmt.Println("-----------------------------")

	return nil
}

func UpdateExpense(expense models.Expense) error {
	if err := database.UpdateInExpenseTable(expense); err != nil {
		return err
	}

	if expense.Members != nil && len(expense.Members) > 0 {
		err := updateRelatedTables(expense)
		if err != nil {
			return err
		}
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
