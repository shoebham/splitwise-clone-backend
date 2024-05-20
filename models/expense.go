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
	Eid         int             `json:"Eid" sql:"eid"`
	Description string          `json:"Description" sql:"description"`
	Amount      int             `json:"Amount" sql:"amount"`
	UserAdded   int             `json:"User_added" sql:"added_by"`
	UserPaid    int             `json:"User_paid" sql:"paid_by"`
	Members     map[int]float64 `json:"Members" sql:"members"`
	IsEqually   bool            `json:"IsEqually" sql:"isequal"`
	IsSettled   bool            `json:"IsSettled" sql:"issettled"`
	Gid         int             `json:"Gid" sql:"gid"`
}
