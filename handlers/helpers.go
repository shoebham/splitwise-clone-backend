package handlers

import (
	"splitwise-backend/models"
	"strconv"
)

var userIds []string
var shares []float64

func updateRelatedTables(expense models.Expense) error {
	err := updateUserAfterCreatingOrUpdatingExpense(expense)
	if err != nil {
		return err
	}
	err = updateBalanceInGroup(expense)
	if err != nil {
		return err
	}
	return nil
}
func updateUserAfterCreatingOrUpdatingExpense(expense models.Expense) error {
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

func updateBalanceInGroup(expense models.Expense) error {

	// for each user get its share and update in group table
	// group.balances [user_paid->{member1,amt...}]
	// group.balances [member1->{user_paid,-amt...}]
	groupId := expense.Gid
	groupIdStr := strconv.Itoa(groupId)
	group := GetGroupById([]string{groupIdStr})

	var groupBalance []models.GroupBalance
	if len(userIds) > 0 {
		for i, uid := range userIds {
			bal := models.GroupBalance{
				uid,
				int(shares[i]),
			}
			groupBalance = append(groupBalance, bal)
		}
	}
	group[0].Balances[expense.UserPaid] = groupBalance
	for i, uid := range userIds {
		uidInt, err := strconv.Atoi(uid)
		if err != nil {
			// Handle the error appropriately, perhaps log it or skip this iteration
			return err
		}
		group[0].Balances[uidInt] = []models.GroupBalance{{UserId: strconv.Itoa(expense.UserPaid), Amount: -int(shares[i])}}
	}
	return nil
}
