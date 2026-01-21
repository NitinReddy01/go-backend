package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	DB_URL string
}

func LoadConfig() *Config {

	if os.Getenv("ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Note: No .env file found or error loading it. Assuming OS environment variables will be used")
		}
	}

	portString := getEnv("PORT")

	_, err := strconv.Atoi(portString)
	if err != nil {
		log.Fatalf("Invalid port: %s", portString)
	}
	if portString == "" || err != nil {
		log.Fatalf("Invalid Port %s:", portString)
	}

	dbUrl := getEnv("PG_URL")

	if dbUrl == "" {
		log.Fatal("Missing Postgres DB url")
	}

	config := &Config{
		Port:   portString,
		DB_URL: dbUrl,
	}

	return config
}

func getEnv(key string) string {
	value, ok := os.LookupEnv(key)

	if ok {
		return value
	}

	return ""
}
