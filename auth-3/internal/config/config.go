package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	AppEnv string
	Dns string
	JwtSecrect string
	JwtDuration string
	SuperAdminEmail string
	SuperAdminPassword string
	AdminEmail string
	AdminPassword string
}

func mustGetEnv(key string) string {
	value := os.Getenv(key)

	if value == "" {
		fmt.Printf("%s is required\n", key)
	}
	return value
}

func LoadEnv() *Config{
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not loaded")
	}

	
	return &Config{
		Port: mustGetEnv("PORT"),
		AppEnv: mustGetEnv("APP_ENV"),
		Dns: mustGetEnv("DNS"),
		JwtSecrect: mustGetEnv("JWT_SECRECT"),
		JwtDuration: mustGetEnv("JWT_DURATION"),
		SuperAdminEmail: mustGetEnv("SUPER_ADMIN_EMAIL"),
		SuperAdminPassword: mustGetEnv("SUPER_ADMIN_PASSWORD"),
		AdminEmail: mustGetEnv("ADMIN_EMAIL"),
		AdminPassword: mustGetEnv("ADMIN_PASSWORD"),
	}
}