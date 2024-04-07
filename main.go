package main

import (
	"fmt"
	"splitwise-backend/app"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Initialize the application
	if err := app.Init(); err != nil {
		fmt.Println("Error initializing application:", err)
		return
	}

	// database.GetAllData("users")
}
