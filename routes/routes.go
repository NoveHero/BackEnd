package routes

import (
	types "NoveHeroAPI/glob"
	// llm "NoveHeroAPI/llm"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
)

func AuthMiddleware(db *gorm.DB) fiber.Handler { // db instance is now passed
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")

		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Missing Authorization header"})
		}

		// Remove "Bearer " prefix if present
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		token, err := jwt.ParseWithClaims(tokenString, &types.Claims{}, func(token *jwt.Token) (interface{}, error) {
			// Replace with your actual secret key
			return []byte("your-jwt-secret"), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
		}

		claims, ok := token.Claims.(*types.Claims)
		if !ok {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Invalid token claims"}) // More specific error
		}

		// Optionally: Check if user exists in the database (important for revoked tokens)
		// var user model.User // Replace model.User with your actual User model
		// if err := db.First(&user, claims.UserID).Error; err != nil {
		//     return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid user"})
		// }

		// Store the claims in the context Locals for use in handlers
		c.Locals("claims", claims)

		return c.Next()
	}
}
func SetupRoutes(app *fiber.App, db *gorm.DB, openaiClient *openai.Client) {
	api := app.Group("/api")

	api.Post("/auth/signup", SignupHandler(db))
	api.Post("/auth/login", LoginHandler(db))
	api.Get("/llm-permission", AuthMiddleware(db), PermitLLMConnection(db))
	api.Post("/auth/change-password", AuthMiddleware(db), ChangePasswordHandler(db))
	api.Get("/me/", AuthMiddleware(db), MeHandler(db))
	api.Patch("/me/change-nickname", AuthMiddleware(db), ChangeNicknameHandler(db))
	api.Get("/me/stories", AuthMiddleware(db), GetStories(db))
	api.Get("/me/stories/:uuid", AuthMiddleware(db), GetStory(db))
	api.Post("/me/stories", AuthMiddleware(db), NewStory(db))
	api.Patch("/me/stories/:uuid/update", AuthMiddleware(db), UpdateStory(db))
	api.Patch("/me/stories/:uuid/change-title", AuthMiddleware(db), ChangeStoryTitle(db))
	api.Delete("/me/stories/:uuid/", AuthMiddleware(db), DeleteStory(db))
	api.Put("/me/stories/meta/:uuid/", AuthMiddleware(db), UpdatMetaData(db))

}
