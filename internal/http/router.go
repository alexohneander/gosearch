package http

import (
	"git.dev-null.rocks/alexohneander/gosearch/internal/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func configureRoutes(app *fiber.App) *fiber.App {
	// Index
	app.Post("/api/index", controller.AddDocumentToIndex)

	// Search
	app.Get("/api/search/:query", controller.SearchQuery)

	// Monitor
	app.Get("/metrics", monitor.New(monitor.Config{Title: "gosearch Metrics"}))

	// Health Checks
	// app.Get(healthcheck.DefaultLivenessEndpoint, healthcheck.NewHealthChecker())
	// app.Get(healthcheck.DefaultReadinessEndpoint, healthcheck.NewHealthChecker())
	// app.Get(healthcheck.DefaultStartupEndpoint, healthcheck.NewHealthChecker())
	return app
}
