package models

/*
Expense {
    ExpenseID,
    Description,
    Amount,
    User Added,
    User Paid,
    Map<User, double> Members,
    isEqually,
    isSettled,
    GroupID
}
*/

type Expense struct {
	ExpenseID   int               `json:"expenseID"`
	Description string            `json:"description"`
	Amount      int               `json:"amount"`
	User_added  string            `json:"user_added"`
	User_paid   string            `json:"user_paid"`
	Members     map[*User]float64 `json:"members"`
	IsEqually   bool              `json:"isEqually"`
	IsSettled   bool              `json:"isSettled"`
	GroupID     int               `json:"groupID"`
}
