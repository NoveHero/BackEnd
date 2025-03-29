package controllers

import (
	database "NoveHeroAPI/database"
	types "NoveHeroAPI/glob"

	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Stories(c *fiber.Ctx, db *gorm.DB) error {
	claims := c.Locals("claims").(*types.Claims)
	var stories []database.Story // Directly query the stories table

	if err := db.Preload("MetaData"). // Preload StoryMetaData
						Where("user_uuid = ?", claims.UserID).
						Select("created_at, updated_at, title, content, uuid").
						Find(&stories).Error; err != nil {
		// Handle error appropriately
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"results": stories,
	})
}
func NewStory(c *fiber.Ctx, db *gorm.DB) error {
	claims := c.Locals("claims").(*types.Claims) // Access claims from middleware
	var user database.User

	if err := db.First(&user, "uuid = ?", claims.UserID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"}) // Or unauthorized
	}

	// Parse the request body to get the title and metadata
	var input struct {
		Title    string `json:"title"`
		MetaData struct {
			Genre        string `json:"genre"`
			Setting      string `json:"setting"`
			Protagonist  string `json:"protagonist"`
			Antagonist   string `json:"antagonist"`
			ConflictInfo string `json:"conflict_info"`
			DialogueInfo string `json:"dialogue_info"`
			Theme        string `json:"theme"`
			Tone         string `json:"tone"`
			Pacing       string `json:"pacing"`
		} `json:"meta_data"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request body"})
	}

	storyUUID := uuid.New().String()
	metaDataUUID := uuid.New().String()

	// Initialize a new story with the title and empty content
	newStory := database.Story{
		Title:    input.Title,
		Content:  "",        // Initialize with empty content
		UserUUID: user.UUID, // Associate the story with the user
		UUID:     storyUUID,
	}

	// Initialize metadata for the story
	newStoryMetaData := database.StoryMetaData{
		UUID:         metaDataUUID,
		Genre:        input.MetaData.Genre,
		Setting:      input.MetaData.Setting,
		Protagonist:  input.MetaData.Protagonist,
		Antagonist:   input.MetaData.Antagonist,
		ConflictInfo: input.MetaData.ConflictInfo,
		DialogueInfo: input.MetaData.DialogueInfo,
		Theme:        input.MetaData.Theme,
		Tone:         input.MetaData.Tone,
		Pacing:       input.MetaData.Pacing,
		StoryUUID:    storyUUID,
	}

	// Begin a transaction
	tx := db.Begin()

	// Save the new story to the database
	if err := tx.Create(&newStory).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create story"})
	}

	// Save the story metadata to the database
	if err := tx.Create(&newStoryMetaData).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create story metadata"})
	}

	// Commit the transaction
	tx.Commit()

	// Return the created story
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Story created successfully",
	})
}

func UpdateStory(c *fiber.Ctx, db *gorm.DB) error {
	claims := c.Locals("claims").(*types.Claims) // Access claims from middleware
	var user database.User

	if err := db.First(&user, "uuid = ?", claims.UserID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"}) // Or unauthorized
	}

	var input struct {
		Content string `json:"content"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request body"})
	}
	storyUUID := c.Params("uuid")
	var story database.Story

	if err := db.Model(&database.Story{}).Where("user_uuid = ?", claims.UserID).Where("uuid = ?", storyUUID).First(&story).Error; err != nil {
		// Handle error appropriately
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "story not found"})
	}
	story.Content = input.Content
	if err := db.Save(&story).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error updating story"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"uuid": story.UUID})

}

func ChangeStoryTitle(c *fiber.Ctx, db *gorm.DB) error {
	claims := c.Locals("claims").(*types.Claims) // Access claims from middleware
	var user database.User

	if err := db.First(&user, "uuid = ?", claims.UserID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"}) // Or unauthorized
	}

	var input struct {
		Title string `json:"title"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request body"})
	}
	storyUUID := c.Params("uuid")
	var story database.Story

	if err := db.Model(&database.Story{}).Where("user_uuid = ?", claims.UserID).Where("uuid = ?", storyUUID).First(&story).Error; err != nil {
		// Handle error appropriately
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "story not found"})
	}
	story.Title = input.Title
	if err := db.Save(&story).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error updating story"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"uuid": story.UUID})

}

func DeleteStory(c *fiber.Ctx, db *gorm.DB) error {
	claims := c.Locals("claims").(*types.Claims) // Access claims from middleware
	var user database.User

	if err := db.First(&user, "uuid = ?", claims.UserID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"}) // Or unauthorized
	}
	storyUUID := c.Params("uuid")
	var story database.Story

	if err := db.Model(&database.Story{}).Where("user_uuid = ?", claims.UserID).Where("uuid = ?", storyUUID).First(&story).Error; err != nil {
		// Handle error appropriately
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "story not found"})
	}

	if err := db.Delete(&story); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Can not Delete Story."})

	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Deleted Successfuly"})
}

func UpdateMetaData(c *fiber.Ctx, db *gorm.DB) error {
	// Extract the story UUID from the URL parameters
	storyUUID := c.Params("uuid")

	// Parse the request body to get new metadata values
	var input struct {
		Genre        string `json:"genre"`
		Setting      string `json:"setting"`
		Protagonist  string `json:"protagonist"`
		Antagonist   string `json:"antagonist"`
		ConflictInfo string `json:"conflict_info"`
		DialogueInfo string `json:"dialogue_info"`
		Theme        string `json:"theme"`
		Tone         string `json:"tone"`
		Pacing       string `json:"pacing"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request body"})
	}

	// Find the existing StoryMetaData record using the Story UUID
	var metaData database.StoryMetaData
	if err := db.Where("story_uuid = ?", storyUUID).First(&metaData).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Metadata not found"})
	}

	// Update the StoryMetaData with new values
	metaData.Genre = input.Genre
	metaData.Setting = input.Setting
	metaData.Protagonist = input.Protagonist
	metaData.Antagonist = input.Antagonist
	metaData.ConflictInfo = input.ConflictInfo
	metaData.DialogueInfo = input.DialogueInfo
	metaData.Theme = input.Theme
	metaData.Tone = input.Tone
	metaData.Pacing = input.Pacing

	// Save the updated metadata to the database
	if err := db.Save(&metaData).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update metadata"})
	}

	// Return a success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Metadata updated successfully",
	})
}

func GetStory(c *fiber.Ctx, db *gorm.DB) error {
	// Access claims from middleware, which should contain the User ID (UUID)
	claims := c.Locals("claims").(*types.Claims)

	// Fetch the story UUID from the URL parameters
	storyUUID := c.Params("uuid")

	// Initialize a container for the story, along with its metadata
	var story database.Story

	// Query for the story, preload the associated metadata, and ensure the story belongs to the user
	if err := db.Preload("MetaData").
		Where("uuid = ? AND user_uuid = ?", storyUUID, claims.UserID).
		First(&story).Error; err != nil {
		// Handle the case when the story is not found or does not belong to the user
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Story not found"})
	}

	// Successfully retrieved story and metadata
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"story": story,
	})
}
