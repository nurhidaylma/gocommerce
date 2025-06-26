package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func AdminOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if role, ok := c.Locals("role").(string); !ok || role != "admin" {
			return c.Status(403).JSON(fiber.Map{"error": "admin only"})
		}
		return c.Next()
	}
}

func IsAuthorized(resourceOwnerID uint, userID uint) bool {
	return resourceOwnerID == userID
}
