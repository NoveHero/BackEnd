package glob

import (
	"database/sql/driver"
	"errors"

	"github.com/golang-jwt/jwt/v4" // Updated JWT package
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

type LLMModelType string

const (
	GPT4o            LLMModelType = "gpt4o-mini"
	GeminiFlash1Dot5 LLMModelType = "gemini-1.5-flash"
)

// Implement String method for ModelType to satisfy fmt.Stringer interface
func (m LLMModelType) String() string {
	return string(m)
}

// Implement the sql.Scanner interface for ModelType
func (m *LLMModelType) Scan(value interface{}) error {
	if v, ok := value.(string); ok {
		*m = LLMModelType(v)
		return nil
	}
	return errors.New("failed to scan ModelType")
}

// Implement the driver.Valuer interface for ModelType
func (m LLMModelType) Value() (driver.Value, error) {
	return string(m), nil
}
