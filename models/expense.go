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
	ExpenseID   int
	Description string
	Amount      int
	User_added  string
	User_paid   string
	Members     map[*User]int
	isEqually   bool
	isSettled   bool
	GroupID     int
}