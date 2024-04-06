package config

import (
	"log"
	"splitwise-backend/constants"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load()
	constants.InitVars()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return err
	}
	return nil
}
