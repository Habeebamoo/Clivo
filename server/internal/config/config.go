package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func Initialize() {
	godotenv.Load()
}

func Get(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("%s does not exists", strings.ToLower(key))
	}
	return value, nil
}