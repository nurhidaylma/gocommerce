package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			return c.Status(401).JSON(fiber.Map{"error": "missing or invalid token"})
		}

		tokenStr := strings.TrimPrefix(auth, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{"error": "invalid or expired token"})
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Locals("user_id", uint(claims["user_id"].(float64)))
		c.Locals("role", claims["role"].(string))
		return c.Next()
	}
}
