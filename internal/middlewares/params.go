package middlewares

import (
	// Your custom logger

	"fiber-boilerplate/pkg"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
)

// Params sets up middlewares for the application, including logging with zap.
func Params(a *fiber.App) {
	// Use custom zap logger middleware to log incoming requests
	a.Use(func(c *fiber.Ctx) error {
		// Log the incoming request with zap
		pkg.Logger.Info("Incoming request", zap.String("method", c.Method()), zap.String("path", c.Path()))

		// Continue to the next middleware/handler
		return c.Next()
	})

	// Use CORS middleware with credentials allowed
	a.Use(cors.New(cors.Config{
		AllowOrigins:     "*",                                      // Permite todos los orígenes
		AllowCredentials: false,                                    // Deshabilita el uso de credenciales
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS", // Métodos permitidos
	}))

	// Use recovery middleware to handle panics
	a.Use(recover.New())
}
