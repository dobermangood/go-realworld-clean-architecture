package httpserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func New() *fiber.App {
	app := fiber.New()

	app.Use(
		cors.New(),
		recover.New(),
		logger.New(),
	)

	return app
}
