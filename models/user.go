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
	UserID  int
	Name    string
	Balance int
	Owes    map[*User]float64
	Owed    map[*User]float64
}
