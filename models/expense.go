package models

/*
Expense {
    Eid,
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
	Eid         int             `json:"Eid"`
	Description string          `json:"Description"`
	Amount      int             `json:"Amount"`
	User_added  string          `json:"User_added"`
	User_paid   string          `json:"User_paid"`
	Members     map[int]float64 `json:"Members"`
	IsEqually   bool            `json:"IsEqually"`
	IsSettled   bool            `json:"IsSettled"`
	Gid         int             `json:"Gid"`
}
