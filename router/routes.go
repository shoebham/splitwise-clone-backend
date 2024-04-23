package router

import (
	"splitwise-backend/models"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	SetupUserRoutes(app)
	SetupGroupRoutes(app)
	SetupExpenseRoutes(app)
}

var fakeUsers []models.User

func checkId(c fiber.Ctx) (int, error) {
	idStr := c.Params("id")
	idInt, idErr := strconv.Atoi(idStr)
	if idErr != nil {
		return -1, c.Status(400).JSON(fiber.Map{
			"message": "Invalid id",
		})
	}
	return idInt, nil
}
