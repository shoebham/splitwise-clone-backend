package handlers

import (
	"splitwise-backend/database"

	"github.com/gofiber/fiber/v3"
)

func GetAllGroups(c *fiber.App) {
	database.GetAllData("groups")
}

func GetAllUsers(c *fiber.App) {
	database.SelectFromUsers()
}
