package constants

import "os"

var DB_USER, DB_PASS string

func InitVars() {
	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
}
