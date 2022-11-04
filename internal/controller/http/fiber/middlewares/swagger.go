package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func RegisterSwagger(app *fiber.App, swaggerUrl string) {
	app.Use("/swagger/*", swagger.New(swagger.Config{
		DeepLinking: false,
		URL:         swaggerUrl,
	}))
}
