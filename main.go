package main

import (
	"fiber-boilerplate/config"
	"fiber-boilerplate/databases"
	"fiber-boilerplate/internal/middlewares"
	"fiber-boilerplate/pkg"
	"fiber-boilerplate/routes"
	"fiber-boilerplate/utils"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-module/carbon/v2"
	"go.uber.org/zap"
)

func main() {
	// Initialize the logger
	pkg.InitLogger()
	defer pkg.Logger.Sync() // Ensure logger flushes any pending logs on exit

	// Log application start
	pkg.Logger.Info("Starting application", zap.String("timestamp", carbon.Now().ToDateTimeString()))

	carbon.Now().ToDateTimeString()
	// Define Fiber config.
	config := config.FiberConfig()

	//start fiber
	app := fiber.New(config)

	// Middlewares.
	pkg.Logger.Info("Applying middlewares...")
	middlewares.Params(app)

	//connect database
	pkg.Logger.Info("Connecting to the database...")
	databases.Connect()

	pkg.Logger.Info("Setting up routes...")
	routes.SetupRoutes(app)
	routes.NotFoundRoute(app) // Register route for 404 Error.

	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		pkg.Logger.Info("Starting server in development mode...")
		utils.StartServer(app)
	} else {
		pkg.Logger.Info("Starting server with graceful shutdown...")
		utils.StartServerWithGracefulShutdown(app)
	}

}
