package config

import (
	"log"
	"os"
	"strings"

	"github.com/NitinReddy01/go-backend/internal/validation"
	"github.com/joho/godotenv"
)

type Config struct {
	Port               string   `validate:"required,numeric"`
	DB_URL             string   `validate:"required,url"`
	CORSAllowedOrigins []string `validate:"required,min=1,dive,required"`
}

func LoadConfig() *Config {

	if err := godotenv.Load(); err != nil {
		log.Fatal("error while loading .env", err)
	}

	rawOrigins := getEnv("CORS_ALLOWED_ORIGINS")

	if rawOrigins == "" {
		log.Fatal("missing allowed CORS origings")
	}

	config := &Config{
		Port:               getEnv("PORT"),
		DB_URL:             getEnv("DB_URL"),
		CORSAllowedOrigins: strings.Split(rawOrigins, ","),
	}

	if err := validation.Validate.Struct(config); err != nil {
		log.Fatal("Invalid cors", err)
	}

	return config
}

func getEnv(key string) string {
	val := strings.TrimSpace(os.Getenv(key))
	if val == "" {
		log.Fatalf("missing env var: %s", key)
	}
	return val
}
