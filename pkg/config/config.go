package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}
	return &Config{
		Port: os.Getenv("PORT"),
	}
}

