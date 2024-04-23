package router

import (
	"fmt"
	"splitwise-backend/handlers"
	"splitwise-backend/models"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	setupUserRoutes(app)
	setupGroupRoutes(app)
	setupExpenseRoutes(app)
}

var fakeUsers []models.User

func setupUserRoutes(app *fiber.App) {
	users := app.Group("/user")

	usersArr := createFakeUsers()
	users.Get("/", func(c fiber.Ctx) error {
		fakeUsers = handlers.GetAllUsers(app)
		return c.SendString("Hello Users")

	})
	users.Post("/", func(c fiber.Ctx) error {
		for _, user := range usersArr {
			fmt.Printf("User created\n Name:%s Phone:%s\n", user.Name, user.Number)
			handlers.CreateUser(app, user)
		}
		return nil
	})
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
		groupsArr := createFakeGroups()
		for _, group := range groupsArr {
			handlers.CreateGroup(app, group)
		}
		return c.Status(200).JSON(fiber.Map{
			"message": "Success",
		})
		return nil
	})
	// update group with group id
	groups.Put("/:id", func(c fiber.Ctx) error {

		_, idErr := checkId(c)
		if idErr != nil {
			return idErr
		}
		var updatedGroup models.Group
		if err := c.Bind().Body(&updatedGroup); err != nil {
			return err
		}
		handlers.UpdateGroup(app, updatedGroup)
		return nil

	})
	// delete group with group id
	groups.Delete("/:id", func(c fiber.Ctx) error {

		idInt, idErr := checkId(c)
		if idErr != nil {
			return idErr
		}

		handlers.DeleteGroup(idInt)
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
