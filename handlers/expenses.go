package handlers

import (
	"fmt"
	"splitwise-backend/database"
	"splitwise-backend/models"
	"strconv"
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

	var userIds []string
	var shares []float64
	// get user ids and their shares
	for uid, share := range expense.Members {
		userIds = append(userIds, strconv.Itoa(uid))
		shares = append(shares, share)
	}
	// get all users by id
	users := GetUserById(userIds)
	// update the share of users
	for i, user := range users {
		if user.Owes == nil {
			user.Owes = make(map[int]float64)
		}
		if val, ok := user.Owes[expense.UserPaid]; ok {
			user.Owes[expense.UserPaid] = val + shares[i]
		} else {
			user.Owes[expense.UserPaid] = shares[i]
		}
		users[i] = user
	}

	// update in DB
	for _, user := range users {
		if err := UpdateUser(user); err != nil {
			return err
		}
	}
	fmt.Println("-----------------------------")

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
