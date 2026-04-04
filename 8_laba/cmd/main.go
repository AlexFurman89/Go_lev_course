package main

import (
	"log"

	"8_laba/internal/db"
	"8_laba/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	app.Get("/contacts", func(c *fiber.Ctx) error {
		return handlers.GetContacts(c, database)
	})
	app.Get("/contacts/:id", func(c *fiber.Ctx) error {
		return handlers.GetContactById(c, database)
	})
	app.Post("/contacts", func(c *fiber.Ctx) error {
		return handlers.CreateContact(c, database)
	})
	app.Put("/contacts/:id", func(c *fiber.Ctx) error {
		return handlers.UpdateContact(c, database)
	})
	app.Delete("/contacts/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteContact(c, database)
	})

	app.Listen(":3000")
}
