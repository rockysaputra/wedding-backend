package router

import (
	"wedding-api/handler"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/healthchecker", handler.Hello)
}
