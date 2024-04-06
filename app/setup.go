package app

import "splitwise-backend/config"

func Init() error {
	err := config.LoadEnv()
	if err != nil {
		return err
	}
	return nil
}
