package models

/*
	Group {
	    GroupID,
	    GroupName,
	    List<User> Members
	}
*/
type Group struct {
	Gid       int    `json:"Gid" sql:"gid"`
	GroupName string `json:"GroupName" sql:"group_name"`
	Members   []int  `json:"Members" sql:"members"`
}

// TODO
// add balances column in group user->{owes/owed to,
//+- share} where user owes or is owed to another user indicated by + or
//- balance
// for ex. 1->{2,-100}
// this means user 1 owes 100 to user 2 in this particular group
// remove isSettled from expense
