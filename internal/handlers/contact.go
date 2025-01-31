package handlers

import (
	"fiber-boilerplate/databases"
	"fiber-boilerplate/internal/models"
	"fiber-boilerplate/pkg"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func GetContacts(c *fiber.Ctx) error {
	var contacts []models.Contact
	if err := databases.Database.Find(&contacts).Error; err != nil {
		pkg.Logger.Error("Failed to retrieve contacts", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to retrieve contacts",
		})
	}
	pkg.Logger.Info("Retrieved contacts successfully", zap.Int("count", len(contacts)))
	return c.Status(fiber.StatusOK).JSON(contacts)
}

func GetContact(c *fiber.Ctx) error {
	id := c.Params("id")
	var contact models.Contact

	if err := databases.Database.First(&contact, id).Error; err != nil {
		pkg.Logger.Warn("contact not found", zap.String("id", id), zap.Error(err))
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "contact not found",
		})
	}

	pkg.Logger.Info("Retrieved contact successfully", zap.String("id", id))

	return c.Status(fiber.StatusOK).JSON(contact)
}

func AddContact(c *fiber.Ctx) error {
	contact := new(models.Contact)

	// Parse request body
	if err := c.BodyParser(contact); err != nil {
		pkg.Logger.Error("Invalid request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Save to database
	if err := databases.Database.Create(&contact).Error; err != nil {
		pkg.Logger.Error("Failed to create contact", zap.Any("contact", contact), zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to create contact",
		})
	}

	pkg.Logger.Info("contact created successfully", zap.Any("contact", contact))
	return c.Status(fiber.StatusCreated).JSON(contact)
}

func UpdateContact(c *fiber.Ctx) error {
	id := c.Params("id")
	var contact models.Contact

	// Check if the contact exists
	if err := databases.Database.First(&contact, id).Error; err != nil {
		pkg.Logger.Warn("contact not found for update", zap.String("id", id), zap.Error(err))
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "contact not found",
		})
	}

	// Parse request body
	updateData := new(models.Contact)
	if err := c.BodyParser(updateData); err != nil {
		pkg.Logger.Error("Invalid request body for update", zap.String("id", id), zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Update the contact in the database
	if err := databases.Database.Model(&contact).Updates(updateData).Error; err != nil {
		pkg.Logger.Error("Failed to update contact", zap.String("id", id), zap.Any("updateData", updateData), zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to update contact",
		})
	}

	pkg.Logger.Info("contact updated successfully", zap.String("id", id), zap.Any("updateData", updateData))
	return c.Status(fiber.StatusOK).JSON(contact)
}

func RemoveContact(c *fiber.Ctx) error {
	id := c.Params("id")

	// Delete the contact
	if err := databases.Database.Delete(&models.Contact{}, id).Error; err != nil {
		pkg.Logger.Warn("Failed to delete contact", zap.String("id", id), zap.Error(err))
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "contact not found",
		})
	}

	pkg.Logger.Info("contact deleted successfully", zap.String("id", id))
	return c.SendStatus(fiber.StatusOK)
}
