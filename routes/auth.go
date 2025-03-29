package routes

import (
	"NoveHeroAPI/controllers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// const jwtSecret = "your-jwt-secret"

func SignupHandler(db *gorm.DB) fiber.Handler {
	// ... (signup logic from previous example - use db as needed) ...
	return func(c *fiber.Ctx) error {
		return controllers.Signup(c, db)
	}
}

func LoginHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return controllers.Login(c, db)
	}
}

func ChangePasswordHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return controllers.ChangePassword(c, db)
	}
}
