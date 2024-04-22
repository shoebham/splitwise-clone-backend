package models

/*
	Group {
	    GroupID,
	    GroupName,
	    List<User> Members
	}
*/
type Group struct {
	Gid       int    `json:"Gid"`
	GroupName string `json:"GroupName"`
	Members   []User `json:"Members"`
}
