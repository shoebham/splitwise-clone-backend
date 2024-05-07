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
	// for each user update in user table
	// if not equal then iterate over user members, check the sum of share it should be equal to amount
	// for each user update users share, if user paid and in members, do nothing for that user

	//for userId, share := range expense.Members {
	//	if userId == expense.User_paid{
	//
	//	}
	//	user := models.User{
	//		uid:userId,
	//	}
	//
	//	if err := database.UpdateUser(user)
	//}
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
