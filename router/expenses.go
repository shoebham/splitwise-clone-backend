package router

import "github.com/gofiber/fiber/v3"

func SetupExpenseRoutes(app *fiber.App) {
	expenses := app.Group("/expense")

	// get group details
	expenses.Get("/", func(c fiber.Ctx) error {
		return nil
	})
	// create new group
	expenses.Post("/", func(c fiber.Ctx) error {
		return nil
	})
	// update group with group id
	expenses.Put("/:id", func(c fiber.Ctx) error {
		return nil
	})
	// delete group with group id
	expenses.Delete("/:id", func(c fiber.Ctx) error {
		return nil
	})
	// delete group with group id
	expenses.Post("/:id/settleUp", func(c fiber.Ctx) error {
		return nil
	})
}
