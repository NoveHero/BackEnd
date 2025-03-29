package routes

import (
	"NoveHeroAPI/controllers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetStories(db *gorm.DB) fiber.Handler {
	// ... (signup logic from previous example - use db as needed) ...
	return func(c *fiber.Ctx) error {
		return controllers.Stories(c, db)
	}
}

func NewStory(db *gorm.DB) fiber.Handler {
	// ... (signup logic from previous example - use db as needed) ...
	return func(c *fiber.Ctx) error {
		return controllers.NewStory(c, db)
	}
}

func UpdateStory(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return controllers.UpdateStory(c, db)
	}
}

func ChangeStoryTitle(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return controllers.ChangeStoryTitle(c, db)
	}
}

func DeleteStory(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return controllers.DeleteStory(c, db)
	}
}

func UpdatMetaData(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return controllers.UpdateMetaData(c, db)
	}
}

func GetStory(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return controllers.GetStory(c, db)
	}
}
