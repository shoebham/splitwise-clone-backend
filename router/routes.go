package router

import (
	"splitwise-backend/handlers"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {

	setupGroupRoutes(app)
	setupExpenseRoutes(app)
}

func setupGroupRoutes(app *fiber.App) {
	groups := app.Group("/group")

	// get group details
	groups.Get("/", func(c fiber.Ctx) error {
		handlers.GetAllGroups(c.App())
		return nil
	})
	// create new group
	groups.Post("/", func(c fiber.Ctx) error {
		return nil
	})
	// update group with group id
	groups.Put("/:id", func(c fiber.Ctx) error {
		return nil
	})
	// delete group with group id
	groups.Delete("/:id", func(c fiber.Ctx) error {
		return nil
	})
	// add group member and get member id in return
	groups.Post("/addMember", func(c fiber.Ctx) error {
		return nil
	})
	// delete group member with member id
	groups.Delete("/deleteMember/:mid", func(c fiber.Ctx) error {
		return nil
	})

	groups.Post("/:id/updateTransactions", func(c fiber.Ctx) error {
		return nil
	})

}

func setupExpenseRoutes(app *fiber.App) {
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
