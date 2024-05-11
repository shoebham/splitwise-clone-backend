package router

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"splitwise-backend/handlers"
)

func SetupUserRoutes(app *fiber.App) {
	users := app.Group("/user")

	getUserDetails(app, users)
	createUser(users)
}

func createUser(users fiber.Router) {
	usersArr := CreateFakeUsers()
	users.Post("/", func(c fiber.Ctx) error {
		for _, user := range usersArr {
			fmt.Printf("User created\n Name:%s Phone:%s\n", user.Name, user.Number)
			if err := handlers.CreateUser(user); err != nil {
				return err
			}
		}
		return nil
	})
}

func getUserDetails(app *fiber.App, users fiber.Router) fiber.Router {
	return users.Get("/", func(c fiber.Ctx) error {
		fakeUsers = handlers.GetAllUsers()
		return c.SendString("Hello Users")

	})
}
