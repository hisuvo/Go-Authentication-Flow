package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT string
	DNS	string
	SESSION_SECRET	string
}

func LoadEnv() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("env not found")
	}

	return Config{
		PORT: os.Getenv("PORT"),
		DNS: os.Getenv("DNS"),
		SESSION_SECRET: os.Getenv("SESSION_SECRET"),
	}
}