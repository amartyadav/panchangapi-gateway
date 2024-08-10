package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("../../.env")
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("db.go: No .env file found, proceeding with environment variables")
		} else {
			log.Printf("db.go: Error loading .env file: %v", err)
		}
	} else {
		log.Println("db.go: .env file loaded successfully")
	}

}
