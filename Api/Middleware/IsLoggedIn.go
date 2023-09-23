package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func IsLoggedIn(c *fiber.Ctx) error {
	if c.Locals("user") == nil {
		return c.Status(401).SendString("Unauthorized - isLoggedIn")
	}

	return c.Next()
}
