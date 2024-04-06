package models

/*
	Group {
	    GroupID,
	    GroupName,
	    List<User> Members
	}
*/
type Group struct {
	GroupId   int    `json:"GroupId"`
	GroupName string `json:"GroupName"`
	Members   []User `json:"Members"`
}
