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
type GroupBalance struct {
	UserId string `json:"user_id"`
	Amount int    `json:"amount"`
}
type Group struct {
	Gid        int                    `json:"Gid" sql:"gid"`
	GroupName  string                 `json:"GroupName" sql:"group_name"`
	Members    []int                  `json:"Members" sql:"members"`
	Balances   map[int][]GroupBalance `json:"Balances" sql:"balances"`
	IsSimplify bool                   `json:"IsSimplify" sql:"is_simplify"`
}
