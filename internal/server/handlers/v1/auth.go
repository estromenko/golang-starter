package v1

import (
	"github.com/gofiber/fiber/v2"
)

func RegistrationHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"registration": "ok",
		})
	}
}
