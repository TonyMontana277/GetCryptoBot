package bot

import (
	"fmt"

	"github.com/joho/godotenv"
)

// LoadConfig loads configuration from a file (e.g., .env)
func LoadConfig() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("Error loading .env file: %v", err)
	}
	return nil
}
