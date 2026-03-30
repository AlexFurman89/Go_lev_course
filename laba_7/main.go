package main

import (
	"laba_7/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/notes", handlers.GetNotes)
	app.Get("/notes/:id", handlers.GetNoteByID)
	app.Post("/notes", handlers.CreateNote)
	app.Put("/notes/:id", handlers.UpdateNote)
	app.Delete("/notes/:id", handlers.DeleteNote)
	app.Listen(":3000")
}
