package models

/*

User {
    UserID,
    Name,
    Balance,
    Map<User, double> Owes,
    Map<User, double> Owed,
}

*/

type User struct {
	UserID  int               `json:"UserID"`
	Name    string            `json:"Name"`
	Balance int               `json:"Balance"`
	Owes    map[*User]float64 `json:"Owes"`
	Owed    map[*User]float64 `json:"Owed"`
}
