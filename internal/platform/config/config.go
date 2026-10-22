package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	Env         string
	DatabaseUrl string
}

func MustLoad() Config {
	// Load .env for local development if available.
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading from system environment")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "production"
	}

	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		panic("DATABASE is required")
	}

	return Config{
		Port:        port,
		Env:         env,
		DatabaseUrl: databaseUrl,
	}
}
