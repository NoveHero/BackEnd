package main

import (
	"log"

	"NoveHeroAPI/database"
	"NoveHeroAPI/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sashabaranov/go-openai"
)

func main() {
	// Initialize database

	openaiClient := openai.NewClient("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")
	db := database.InitDatabase()
	if db == nil {
		log.Fatal("Failed to connect to database")
	}
	defer database.CloseDatabase(db) // Close database connection before exit

	app := fiber.New()
	app.Use(cors.New())
	routes.SetupRoutes(app, db, openaiClient) // Pass the DB instance to routes setup
	log.Fatal(app.Listen(":3000"))
}
