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
	Uid     int               `json:"Uid"`
	Name    string            `json:"Name"`
	Balance int               `json:"Balance"`
	Owes    map[*User]float64 `json:"Owes"`
	Owed    map[*User]float64 `json:"Owed"`
	Number  string            `json:"Number"`
}
