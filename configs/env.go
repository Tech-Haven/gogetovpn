package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		fmt.Println("Env loaded...")
	}
}

func AuthSecret() string {
	return os.Getenv("AUTH_SECRET")
}
