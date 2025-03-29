package controllers

import (
	types "NoveHeroAPI/glob"
	"time"

	"NoveHeroAPI/database"
	"NoveHeroAPI/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Signup(c *fiber.Ctx, db *gorm.DB) error {
	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request body"})
	}

	// Generate a new UUID
	userUUID := uuid.New().String()

	// Hash the password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to hash password"})
	}

	// Create the user with the generated UUID
	if err := db.Create(&database.User{
		UUID:     userUUID,
		Username: input.Username,
		Email:    input.Email,
		Password: hashedPassword,
		Nickname: nil,
	}).Error; err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "User already exists"}) // Handle unique constraint violation
	}

	return c.JSON(fiber.Map{"message": "User created successfully", "uuid": userUUID})
}

// const jwtSecret = "your-jwt-secret" // In production, store this securely

func Login(c *fiber.Ctx, db *gorm.DB) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&input); err != nil { // Corrected BodyParser
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid credentials"}) // Return specific message
	}

	var user database.User
	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	// Create JWT token
	claims := types.Claims{
		UserID: user.UUID,
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte("your-jwt-secret"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{"accessToken": signedToken})
}

func ChangePassword(c *fiber.Ctx, db *gorm.DB) error {
	var input struct {
		OldPassword  string `json:"old"`
		NewPassword1 string `json:"new"`
		NewPassword2 string `json:"repeat"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid input"})
	}

	if input.NewPassword1 != input.NewPassword2 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Passwords do not match"})
	}

	claims := c.Locals("claims").(*types.Claims) // Access claims from middleware
	var user database.User

	if err := db.First(&user, "uuid = ?", claims.UserID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"}) // Or unauthorized
	}

	// Compare old password with stored hash
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(input.OldPassword))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Incorrect old password"})
	}

	// Hash the new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword1), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error hashing password"})
	}

	user.Password = hashedPassword

	if err := db.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error updating password"})
	}

	return c.JSON(fiber.Map{"message": "Password changed successfully"})
}

func ChangeNickname(c *fiber.Ctx, db *gorm.DB) error {
	var input struct {
		Nickname string `json:"nickname"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid input"})
	}

	claims := c.Locals("claims").(*types.Claims) // Access claims from middleware
	var user database.User

	if err := db.First(&user, "uuid = ?", claims.UserID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"}) // Or unauthorized
	}

	// Compare old password with stored hash

	// Hash the new password

	user.Nickname = &input.Nickname

	if err := db.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error updating nickname"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Nickname changed successfully"})
}

func Me(c *fiber.Ctx, db *gorm.DB) error {
	claims := c.Locals("claims").(*types.Claims) // Access claims from middleware
	var user database.User

	if err := db.First(&user, "uuid = ?", claims.UserID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"}) // Or unauthorized
	}

	return c.JSON(fiber.Map{
		"username":   user.Username,
		"nickname":   user.Nickname,
		"email":      user.Email,
		"join":       user.CreatedAt,
		"lastUpdate": user.UpdatedAt,
	})
}

func PermitLLMConnection(c *fiber.Ctx, db *gorm.DB) error {
	claims := c.Locals("claims").(*types.Claims) // Access claims from middleware
	var user database.User

	if err := db.First(&user, "uuid = ?", claims.UserID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"}) // Or unauthorized
	}

	return c.JSON(fiber.Map{
		"permission": user.UUID,
	})

}
