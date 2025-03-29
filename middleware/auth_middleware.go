package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

const jwtSecret = "your-jwt-secret"

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(jwtSecret),
		ErrorHandler: jwtError,
		// Other JWT config options as needed
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	// Handle JWT errors
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "Unauthorized",
	})
}
