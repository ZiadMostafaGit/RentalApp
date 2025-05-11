package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct { // Exported struct
	AppPort    string // Exported fields
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() *Config { // Exported function
	err := godotenv.Load("/home/ziad/git/rental_app/.env")
	if err != nil {
		log.Fatal("error loading the .env file: ", err)
	}
	return &Config{
		AppPort:    getEnv("APP_PORT", "8080"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBUser:     getEnv("DB_USER", ""),
		DBName:     getEnv("DB_NAME", ""),
	}
}

func getEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}
