package handlers

import (
	"splitwise-backend/database"
	"splitwise-backend/models"

	"github.com/gofiber/fiber/v3"
)

func GetAllGroups(c *fiber.App) {
	database.GetAllData("groups")
}

func GetAllUsers(c *fiber.App) []models.User {
	return database.SelectFromUsers()
}

func CreateGroup(c *fiber.App, group models.Group) {
	database.InsertInGroupTable(group)
}
