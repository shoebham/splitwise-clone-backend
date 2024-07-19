package handlers

import (
	"splitwise-backend/database"
	"splitwise-backend/models"

	"github.com/gofiber/fiber/v3"
)

func GetAllGroups(c *fiber.App) {
	err := database.GetAllData("groups")
	if err != nil {
		return
	}
}

func CreateGroup(group models.Group) error {
	err := database.InsertInGroupTable(group)
	return err
}

func UpdateGroup(group models.Group) error {
	err := database.UpdateGroup(group)
	if err != nil {
		return err
	}
	return nil
}

func GetGroupById(id []string) []models.Group {
	return database.SelectFromGroups(id)
}

func DeleteGroup(gid int) error {
	err := database.DeleteGroup(gid)
	if err != nil {
		return err
	}
	return nil
}

func AddMembersToGroup(gid int, memberList []int) error {
	err := database.AddMembersToGroup(gid, memberList)
	if err != nil {
		return err
	}
	return nil
}

func DeleteMembersFromGroup(gid int, memberList []int) error {
	err := database.DeleteMembersFromGroup(gid, memberList)
	if err != nil {
		return err
	}
	return nil
}
