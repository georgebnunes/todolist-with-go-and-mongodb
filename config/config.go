package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI      string
	MongoDatabase string
	ServerPort    string
}

func Load() *Config {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading from environment...")
	}

	return &Config{
		MongoURI:      os.Getenv("MONGO_URI"),
		MongoDatabase: os.Getenv("MONGO_DATABASE"),
		ServerPort:    os.Getenv("SERVER_PORT"),
	}
}
