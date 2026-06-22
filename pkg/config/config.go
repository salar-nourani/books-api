package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	DBURL  string 
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	return &Config{
		Port:  getEnv("PORT", "8080"),   
		DBURL: getEnv("DB_URL", ""),      
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

