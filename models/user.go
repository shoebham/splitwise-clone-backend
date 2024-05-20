package models

/*

User {
    Uid,
    Name,
    Balance,
    Map<User, double> Owes,
    Map<User, double> Owed,
}

*/

type User struct {
	Uid     int             `json:"Uid" sql:"uid"`
	Name    string          `json:"Name" sql:"name"`
	Balance float64         `json:"Balance" sql:"balance"`
	Owes    map[int]float64 `json:"Owes" sql:"owes"`
	Owed    map[int]float64 `json:"Owed" sql:"owed"`
	Number  string          `json:"Number" sql:"number"`
}
