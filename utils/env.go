package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Env() {
	envErr := godotenv.Load("../.env")
	if envErr != nil {
		fmt.Printf("Error loading env: %s", envErr)
		os.Exit(1)
	}
}
