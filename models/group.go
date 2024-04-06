package models

/*
	Group {
	    GroupID,
	    GroupName,
	    List<User> Members
	}
*/
type Group struct {
	GroupId   int
	GroupName string
	Members   []User
}
