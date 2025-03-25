package http

import (
	"git.dev-null.rocks/alexohneander/gosearch/internal/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func configureRoutes(app *fiber.App) *fiber.App {
	// Index
	app.Get("/", controller.Index)
	app.Get("/test", controller.Index)

	// Search
	app.Get("/api/search/:index/:query", controller.SearchQuery)

	// Monitor
	app.Get("/metrics", monitor.New(monitor.Config{Title: "gosearch Metrics"}))

	return app
}
