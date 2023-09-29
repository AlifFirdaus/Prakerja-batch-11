package config

import (
	"os"
	"github.com/joho/godotenv"
)

// LoadEnv is a function that loads environment variables from a .env file.
func main() {
	os.Getenv(".env")
	// The godotenv.Load() function loads the environment variables from the .env file.
	// If there is an error, the function panics with the message "failed to load env file".
	if err := godotenv.Load(); err != nil {
		panic("failed to load env file")
	}
}
