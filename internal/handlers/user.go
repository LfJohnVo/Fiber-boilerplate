package handlers

import (
	"fiber-boilerplate/databases"
	"fiber-boilerplate/internal/models"
	"fiber-boilerplate/pkg"

	_ "fiber-boilerplate/docs" // Importa la documentación generada
	"github.com/gofiber/fiber/v2"
	_ "github.com/swaggo/fiber-swagger"
	"go.uber.org/zap"
)

// GetUsers godoc
// @Summary Get all users
// @Description Retrieve a list of users
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]interface{}
// @Router /api/users [get]
func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	if err := databases.Database.Find(&users).Error; err != nil {
		pkg.Logger.Error("Failed to retrieve users", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"error": "Unable to retrieve users",
		})
	}
	pkg.Logger.Info("Retrieved users successfully", zap.Int("count", len(users)))
	return c.Status(fiber.StatusOK).JSON(users)
}

// GetUser godoc
// @Summary Get a user by ID
// @Description Retrieve a single user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} map[string]interface{} // Usé map[string]interface{} en lugar de fiber.Map
// @Router /api/users/{id} [get]
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	if err := databases.Database.First(&user, id).Error; err != nil {
		pkg.Logger.Warn("User not found", zap.String("id", id), zap.Error(err))
		return c.Status(fiber.StatusNotFound).JSON(map[string]interface{}{
			"error": "User not found",
		})
	}

	pkg.Logger.Info("Retrieved user successfully", zap.String("id", id))

	return c.Status(fiber.StatusOK).JSON(user)
}

// AddUser godoc
// @Summary Create a new user
// @Description Add a new user to the database
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User object"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/users [post]
func AddUser(c *fiber.Ctx) error {
	user := new(models.User)

	// Parse request body
	if err := c.BodyParser(user); err != nil {
		pkg.Logger.Error("Invalid request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"error": "Invalid request body",
		})
	}

	// Save to database
	if err := databases.Database.Create(&user).Error; err != nil {
		pkg.Logger.Error("Failed to create user", zap.Any("user", user), zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"error": "Unable to create user",
		})
	}

	pkg.Logger.Info("User created successfully", zap.Any("user", user))
	return c.Status(fiber.StatusCreated).JSON(user)
}

// UpdateUser godoc
// @Summary Update an existing user
// @Description Update the details of a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body models.User true "Updated user data"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]interface{} // Usé map[string]interface{} en lugar de fiber.Map
// @Failure 404 {object} map[string]interface{} // Usé map[string]interface{} en lugar de fiber.Map
// @Failure 500 {object} map[string]interface{} // Usé map[string]interface{} en lugar de fiber.Map
// @Router /api/users/{id} [put]
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	// Check if the user exists
	if err := databases.Database.First(&user, id).Error; err != nil {
		pkg.Logger.Warn("User not found for update", zap.String("id", id), zap.Error(err))
		return c.Status(fiber.StatusNotFound).JSON(map[string]interface{}{
			"error": "User not found",
		})
	}

	// Parse request body
	updateData := new(models.User)
	if err := c.BodyParser(updateData); err != nil {
		pkg.Logger.Error("Invalid request body for update", zap.String("id", id), zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"error": "Invalid request body",
		})
	}

	// Update the user in the database
	if err := databases.Database.Model(&user).Updates(updateData).Error; err != nil {
		pkg.Logger.Error("Failed to update user", zap.String("id", id), zap.Any("updateData", updateData), zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"error": "Unable to update user",
		})
	}

	pkg.Logger.Info("User updated successfully", zap.String("id", id), zap.Any("updateData", updateData))
	return c.Status(fiber.StatusOK).JSON(user)
}

// RemoveUser godoc
// @Summary Remove a user by ID
// @Description Delete a user from the database by ID
// @Tags users
// @Param id path int true "User ID"
// @Success 200 {string} string "User deleted successfully"
// @Failure 404 {object} map[string]interface{} // Usé map[string]interface{} en lugar de fiber.Map
// @Failure 500 {object} map[string]interface{} // Usé map[string]interface{} en lugar de fiber.Map
// @Router /api/users/{id} [delete]
func RemoveUser(c *fiber.Ctx) error {
	id := c.Params("id")

	// Delete the user
	if err := databases.Database.Delete(&models.User{}, id).Error; err != nil {
		pkg.Logger.Warn("Failed to delete user", zap.String("id", id), zap.Error(err))
		return c.Status(fiber.StatusNotFound).JSON(map[string]interface{}{
			"error": "User not found",
		})
	}

	pkg.Logger.Info("User deleted successfully", zap.String("id", id))
	return c.SendStatus(fiber.StatusOK)
}
