package handlers

import (
	"splitwise-backend/database"
	"splitwise-backend/models"

	"github.com/gofiber/fiber/v3"
)

func GetAllGroups(c *fiber.App) {
	database.GetAllData("groups")
}

func CreateGroup(c *fiber.App, group models.Group) {
	database.InsertInGroupTable(group)
}

func UpdateGroup(group models.Group) error {
	err := database.UpdateGroup(group)
	if err != nil {
		return err
	}
	return nil
}

func DeleteGroup(gid int) error {
	err := database.DeleteGroup(gid)
	if err != nil {
		return err
	}
	return nil
}
