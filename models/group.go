package models

/*
	Group {
	    GroupID,
	    GroupName,
	    List<User> Members
		map<User,Pair<user,Share>> Balances
		IsSimplify boolean
	}
*/
type Balance struct {
	UserId string `json:"user_id"`
	Amount int    `json:"amount"`
}
type Group struct {
	Gid        int               `json:"Gid" sql:"gid"`
	GroupName  string            `json:"GroupName" sql:"group_name"`
	Members    []int             `json:"Members" sql:"members"`
	Balances   map[int][]Balance `json:"Balances" sql:"balances"`
	IsSimplify bool              `json:"IsSimplify" sql:"is_simplify"`
}

// TODO
// add balances column in group user->{owes/owed to, +- share}
//where user owes or is owed to another user indicated by + or
//- balance
// for ex. 1->{2,-100}
// this means user 1 owes 100 to user 2 in this particular group
// remove isSettled from expense
