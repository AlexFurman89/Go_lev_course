package handlers

import (
	"laba_7/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetNotes(c *fiber.Ctx) error {
	return c.JSON(models.Notes)
}

func GetNoteByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("id must be number")
	}

	for _, note := range models.Notes {
		if note.ID == id {
			return c.JSON(note)
		}
	}

	return c.Status(404).SendString("not found")
}

func CreateNote(c *fiber.Ctx) error {
	var newNote models.Note

	if err := c.BodyParser(&newNote); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	models.Notes = append(models.Notes, newNote)
	return c.JSON(newNote)
}

func UpdateNote(c *fiber.Ctx) error {
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
			models.Notes[i] = updatedNote
			models.Notes[i].ID = id
			return c.JSON(models.Notes[i])
		}
	}

	return c.Status(404).JSON(fiber.Map{
		"error": "Note not found",
	})
}

func DeleteNote(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
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
}
