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
