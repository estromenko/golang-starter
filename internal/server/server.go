package server

import (
	"fmt"

	swagger "github.com/arsmn/fiber-swagger/v2"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	v1 "golang-starter/internal/server/handlers/v1"
	"golang-starter/pkg/config"

	_ "golang-starter/docs"
)

func Init() error {
	// TODO
	return nil
}

func Run() error {
	app := route()
	return app.Listen(fmt.Sprintf(":%d", config.GetConfig().Server.Port))
}

func route() *fiber.App {
	app := fiber.New(fiber.Config{})

	app.Use(logger.New(logger.Config{}))
	app.Use(cors.New(cors.Config{}))

	public := app.Group("/")
	{
		v1handlers := public.Group("/v1")
		{
			auth := v1handlers.Group("/auth")
			{
				auth.Get("/reg", v1.RegistrationHandler())
			}
		}
	}

	private := app.Group("/")
	{
		private.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("pong")
		})
		private.Get("/swagger/*", swagger.Handler)
	}

	return app
}
