package router

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"splitwise-backend/handlers"
	"splitwise-backend/models"
)

func SetupUserRoutes(app *fiber.App) {
	users := app.Group("/user")

	getUserDetails(app, users)
	createUser(users)
}

func createUser(users fiber.Router) {
	users.Post("/", func(c fiber.Ctx) error {
		var user models.User
		if err := c.Bind().Body(&user); err != nil {
			return err
		}

		if err := handlers.CreateUser(user); err != nil {
			return InternalError(c, err)
		}

		fmt.Printf("User created\n Name:%s Phone:%s\n", user.Name, user.Number)

		return SuccessfulRequest(c, "User Created")
	})
}

func getUserDetails(app *fiber.App, users fiber.Router) fiber.Router {
	return users.Get("/", func(c fiber.Ctx) error {
		fakeUsers = handlers.GetAllUsers()
		return c.SendString("Hello Users")

	})
}
