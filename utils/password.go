package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) ([]byte, error) { // Exported function name
	// ... (password hashing logic from before)
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
