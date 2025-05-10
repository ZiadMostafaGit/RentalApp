package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type config struct {
	app_port    string
	db_host     string
	db_port     string
	db_user     string
	db_password string
	db_name     string
}

func load_config() *config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading the .env file ")
	}
	return &config{
		app_port:    get_env("APP_PORT", "8080"),
		db_host:     get_env("DB_HOST", "localhost"),
		db_port:     get_env("DB_PORT", "3306"),
		db_password: get_env("DB_PASSWORD", ""),
		db_user:     get_env("DB_USER", ""),
		db_name:     get_env("DB_NAME", ""),
	}
}

func get_env(key, default_val string) string {
	val := os.Getenv(key)
	if val == "" {
		return default_val

	}
	return val
}
