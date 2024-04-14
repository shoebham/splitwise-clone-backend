package router

import (
	"fmt"
	"splitwise-backend/handlers"
	"splitwise-backend/models"

	"github.com/go-faker/faker/v4"
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

	usersArr := []models.User{}
	for i := 0; i < 10; i++ {
		name := faker.Name()
		// owes := make(map[*models.User]float64) //rand.Intn(500)
		// owed := make(map[*models.User]float64) //rand.Intn(500)
		// balance := 0                           //owed - owes
		number := faker.Phonenumber()
		usersArr = append(usersArr, models.User{
			Name:   name,
			Number: number,
		})
	}
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
		groupsArr := []models.Group{}
		for i := 0; i < 10; i++ {
			group_name := faker.Word()

			groupsArr = append(groupsArr, models.Group{
				GroupName: group_name,
				Members:   fakeUsers,
			})
		}
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
