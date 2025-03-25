package http

import (
	"git.dev-null.rocks/alexohneander/gosearch/internal/controller"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
)

func configureRoutes(app *fiber.App) *fiber.App {
	// Index
	app.Get("/", controller.Index)
	app.Get("/test", controller.Index)

	// Search
	app.Get("/api/search/:query", controller.SearchQuery)

	// Monitor
	// app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	// Health Checks
	app.Get(healthcheck.DefaultLivenessEndpoint, healthcheck.NewHealthChecker())
	app.Get(healthcheck.DefaultReadinessEndpoint, healthcheck.NewHealthChecker())
	app.Get(healthcheck.DefaultStartupEndpoint, healthcheck.NewHealthChecker())
	return app
}
