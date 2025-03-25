package http

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func StartService() {
	app := fiber.New()

	// Add Logger
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// Configure Routes
	app = configureRoutes(app)

	app.Listen(":3000")
}
