package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseUrl string
	Port string
}

func LoadEnv() Config {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// port, _ := strconv.Atoi(os.Getenv("PORT"))

	return Config{
		DatabaseUrl: os.Getenv("DATABASE_URL"),
		Port: os.Getenv("PORT"),
	}
}