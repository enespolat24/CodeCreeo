package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// png, err := qrcode.Encode("https://www.google.com", qrcode.Medium, 256)

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

}
