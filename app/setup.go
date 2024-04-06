package app

import (
	"log"
	"splitwise-backend/config"

	"github.com/gofiber/fiber/v3"
)

func Init() error {
	err := config.LoadEnv()
	if err != nil {
		return err
	}
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello World")
	})
	log.Fatal(app.Listen(":3000"))
	return nil
}
