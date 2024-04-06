package constants

import "os"

var API_KEY, API_URL string

func InitVars() {
	API_KEY = os.Getenv("API_KEY")
	API_URL = os.Getenv("API_URL")
}
