package handlers

import (
	"splitwise-backend/database"
	"splitwise-backend/models"

	"github.com/gofiber/fiber/v3"
)

func CreateUser(c *fiber.App, user models.User) {
	database.InsertInUserTable(user)
}
