package main

import (
	"laba_7/models"

	"github.com/gofiber/fiber/v2"

	"strconv"
)

func main() {
	app := fiber.New()

	app.Get("/notes", func(c *fiber.Ctx) error {
		return c.JSON(models.Notes)
	})

	app.Get("/notes/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("id must be number")
		}

		for _, note := range models.Notes {
			if note.ID == id {
				data, err := note.MarshalJSON()
				if err != nil {
					return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
				}

				return c.Send(data) // ✅ правильно
			}
		}

		return c.Status(fiber.StatusNotFound).SendString("not found")
	})

	app.Post("/notes", func(c *fiber.Ctx) error {
		var newNote models.Note
		if err := c.BodyParser(&newNote); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}
		models.Notes = append(models.Notes, newNote)

		return c.JSON(newNote)
	})

	app.Put("/notes/:id", func(c *fiber.Ctx) error {
		var updatedNote models.Note

		if err := c.BodyParser(&updatedNote); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid note id",
			})
		}

		for i, note := range models.Notes {
			if note.ID == id {
				models.Notes[i].Title = updatedNote.Title
				models.Notes[i].Content = updatedNote.Content

				return c.JSON(models.Notes[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{
			"error": "Note not found",
		})
	})

	app.Delete("/notes/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid note id",
			})
		}

		for i, note := range models.Notes {
			if note.ID == id {
				models.Notes = append(models.Notes[:i], models.Notes[i+1:]...)
				return c.SendStatus(204)
			}
		}

		return c.Status(404).JSON(fiber.Map{
			"error": "Note not found",
		})
	})
	app.Listen(":3000")
}
