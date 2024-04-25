package router

import (
	"github.com/gofiber/fiber/v3"
	"splitwise-backend/models"
	"strconv"
)

func SetupRoutes(app *fiber.App) {
	SetupUserRoutes(app)
	SetupGroupRoutes(app)
	SetupExpenseRoutes(app)
}

var fakeUsers []models.User

func CheckId(c fiber.Ctx) (int, error) {
	idStr := c.Params("id")
	idInt, idErr := strconv.Atoi(idStr)
	if idErr != nil {
		return -1, c.Status(400).JSON(fiber.Map{
			"message": "Invalid id",
		})
	}
	return idInt, nil
}
func SuccessfulRequest(c fiber.Ctx, message string) error {
	return c.Status(200).JSON(fiber.Map{
		"message": message,
	})

}
func InternalError(c fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error":   true,
		"message": err.Error(),
	})
}
