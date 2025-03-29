package routes

import (
	"NoveHeroAPI/controllers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ChangeNicknameHandler(db *gorm.DB) fiber.Handler {
	// ... (signup logic from previous example - use db as needed) ...
	return func(c *fiber.Ctx) error {
		return controllers.ChangeNickname(c, db)
	}
}
func MeHandler(db *gorm.DB) fiber.Handler {
	// ... (signup logic from previous example - use db as needed) ...
	return func(c *fiber.Ctx) error {
		return controllers.Me(c, db)
	}
}

func PermitLLMConnection(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return controllers.PermitLLMConnection(c, db)
	}
}
