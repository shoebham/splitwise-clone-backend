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

	err := updateUserAfterCreatingOrUpdatingExpense(expense)
	if err != nil {
		return err
	}
	fmt.Println("-----------------------------")

	return nil
}

func updateUserAfterCreatingOrUpdatingExpense(expense models.Expense) error {
	var userIds []string
	var shares []float64
	// get user ids and their shares
	for uid, share := range expense.Members {
		userIds = append(userIds, strconv.Itoa(uid))
		shares = append(shares, share)
	}
	userIds = append(userIds, strconv.Itoa(expense.UserPaid))
	// get all users by id
	users := GetUserById(userIds)

	var userPaid models.User
	// update the share of users
	for i, user := range users {
		if user.Owes == nil {
			user.Owes = make(map[int]float64)
		}
		// user paid
		if expense.UserPaid == user.Uid {
			userPaid = user
			continue
		}
		if val, ok := user.Owes[expense.UserPaid]; ok {
			user.Owes[expense.UserPaid] = val + shares[i]
		} else {
			user.Owes[expense.UserPaid] = shares[i]
		}
		users[i] = user
	}

	// update the user paid with owed amount
	for i, user := range users {
		if userPaid.Owed == nil {
			userPaid.Owed = make(map[int]float64)
		}
		if val, ok := userPaid.Owed[expense.UserPaid]; ok {
			userPaid.Owed[user.Uid] = val + shares[i]
		} else {
			userPaid.Owed[user.Uid] = shares[i]
		}
	}

	// update in DB
	for _, user := range users {
		if err := UpdateUser(user); err != nil {
			return err
		}
	}
	return nil
}

func UpdateExpense(expense models.Expense) error {
	if err := database.UpdateInExpenseTable(expense); err != nil {
		return err
	}

	if expense.Members != nil && len(expense.Members) > 0 {
		err := updateUserAfterCreatingOrUpdatingExpense(expense)
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
