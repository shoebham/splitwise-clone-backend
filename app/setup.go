package app

import (
	"log"
	"splitwise-backend/config"
	"splitwise-backend/database"
	"splitwise-backend/router"

	"github.com/gofiber/fiber/v3"
)

func Init() error {
	err := config.LoadEnv()
	if err != nil {
		return err
	}
	app := fiber.New()
	database.InitDb()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello World")
	})
	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
	return nil
}
