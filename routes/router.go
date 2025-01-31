package routes

import (
	_ "fiber-boilerplate/docs" // Importa la documentación generada
	"fiber-boilerplate/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/swaggo/fiber-swagger"
)

// @title landing-contact
// @version 1.0
// @description API for leads of MKT
// @host localhost:3000
// @BasePath /
func SetupRoutes(app *fiber.App) {
	app.Static("/docs", "./docs")
	app.Get("/docs/*", fiberSwagger.WrapHandler)

	app.Get("/", handlers.Welcome)

	// Middleware
	api := app.Group("/api")

	//Index endpoint
	api.Post("/register", handlers.Register)
	api.Post("/login", handlers.Login)
	api.Get("/authenticated", handlers.AuthenticatedUser)
	api.Post("/logout", handlers.Logout)
	api.Post("/forgot", handlers.Forgot)
	api.Post("/reset", handlers.ResetPassword)
	api.Get("/metrics", monitor.New(monitor.Config{Title: "Metrics"}))

	//Users
	api.Get("/users", handlers.GetUsers)
	api.Get("/users/:id", handlers.GetUser)
	api.Post("/users", handlers.AddUser)
	api.Put("/users/:id", handlers.UpdateUser)
	api.Delete("/users/:id", handlers.RemoveUser)

	//Contacts
	api.Get("/contacts", handlers.GetContacts)
	api.Get("/contacts/:id", handlers.GetContact)
	api.Post("/contacts", handlers.AddContact)
	api.Put("/contacts/:id", handlers.UpdateContact)
	api.Delete("/contacts/:id", handlers.RemoveContact)
}
