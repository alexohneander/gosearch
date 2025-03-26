package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func StartService() {
	app := fiber.New()

	// Add Logger
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	log.SetLevel(log.LevelInfo)

	// Configure Routes
	app = configureRoutes(app)

	app.Listen(":3000")
}
