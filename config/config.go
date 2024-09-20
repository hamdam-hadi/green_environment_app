package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	JWTSecret  string
	AppPort    string
}

func LoadConfig() Config {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config := Config{
		DBUser:     os.Getenv("root"),
		DBPassword: os.Getenv(""),
		DBHost:     os.Getenv("localhost"),
		DBPort:     os.Getenv("3306"),
		DBName:     os.Getenv("green_environment_app"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
		AppPort:    os.Getenv("2024"),
	}

	return config
}
