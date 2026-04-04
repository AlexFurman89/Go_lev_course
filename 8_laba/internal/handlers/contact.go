package handlers

import (
	"8_laba/internal/models"
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetContacts(c *fiber.Ctx, db *sql.DB) error {
	rows, err := db.Query("SELECT id, name, phone FROM contacts")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer rows.Close()

	var contacts []models.Contact

	for rows.Next() {
		var contact models.Contact

		err := rows.Scan(&contact.ID, &contact.Name, &contact.Phone)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		contacts = append(contacts, contact)
	}

	if err := rows.Err(); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(contacts)
}
func GetContactById(c *fiber.Ctx, db *sql.DB) error {

	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid id",
		})
	}

	var contact models.Contact

	err = db.QueryRow(
		"SELECT id, name, phone FROM contacts WHERE id = $1",
		id,
	).Scan(&contact.ID, &contact.Name, &contact.Phone)

	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(404).JSON(fiber.Map{
				"error": "Contact not found",
			})
		}

		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(contact)
}
func CreateContact(c *fiber.Ctx, db *sql.DB) error {
	var newContact models.Contact

	if err := c.BodyParser(&newContact); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := db.QueryRow(
		"INSERT INTO contacts (name, phone) VALUES ($1, $2) RETURNING id",
		newContact.Name,
		newContact.Phone,
	).Scan(&newContact.ID)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(newContact)
}
func UpdateContact(c *fiber.Ctx, db *sql.DB) error {

	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid id",
		})
	}

	var updatedContact models.Contact

	if err := c.BodyParser(&updatedContact); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	result, err := db.Exec(
		"UPDATE contacts SET name = $1, phone = $2 WHERE id = $3",
		updatedContact.Name,
		updatedContact.Phone,
		id,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if rowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": "Contact not found",
		})
	}

	updatedContact.ID = id
	return c.JSON(updatedContact)
}
func DeleteContact(c *fiber.Ctx, db *sql.DB) error {
	id := c.Params("id")

	result, err := db.Exec("DELETE FROM contacts WHERE id = $1", id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "not found"})
	}

	return c.JSON(fiber.Map{"message": "deleted"})
}
