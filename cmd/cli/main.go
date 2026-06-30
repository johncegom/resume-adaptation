package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func setupEnv() {
	// Load environment variables from .env file if present
	_ = godotenv.Load()
}

func main() {
	setupEnv()

	fmt.Println("Resume Adaptation CLI")
	os.Exit(0)
}
