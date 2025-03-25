package http

import (
	"git.dev-null.rocks/alexohneander/gosearch/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func configureRoutes(app *fiber.App) *fiber.App {
	// Index
	app.Get("/", controller.Index)
	app.Get("/test", controller.Index)

	return app
}
