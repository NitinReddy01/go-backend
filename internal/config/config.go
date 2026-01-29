package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/NitinReddy01/go-backend/internal/validation"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTP HTTPConfig
	DB   DBConfig

	CORSAllowedOrigins []string `validate:"required,min=1,dive,required"`
}

type HTTPConfig struct {
	Port         string `validate:"required,numeric"`
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

type DBConfig struct {
	URL string `validate:"required"`

	MaxConns        int32
	MaxConnLifetime time.Duration
	MaxConnIdleTime time.Duration
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, relying on env")
	}

	cfg := &Config{
		HTTP: HTTPConfig{
			Port:         mustEnv("PORT"),
			ReadTimeout:  mustDuration("HTTP_READ_TIMEOUT", 10*time.Second),
			WriteTimeout: mustDuration("HTTP_WRITE_TIMEOUT", 30*time.Second),
			IdleTimeout:  mustDuration("HTTP_IDLE_TIMEOUT", 60*time.Second),
		},
		DB: DBConfig{
			URL:             mustEnv("DB_URL"),
			MaxConns:        mustInt32("DB_MAX_CONNS", 20),
			MaxConnLifetime: mustDuration("DB_MAX_CONN_LIFETIME", time.Hour),
			MaxConnIdleTime: mustDuration("DB_MAX_CONN_IDLE_TIME", 30*time.Minute),
		},
		CORSAllowedOrigins: parseCSV(mustEnv("CORS_ALLOWED_ORIGINS")),
	}

	if err := validation.Validate.Struct(cfg); err != nil {
		log.Fatal("invalid configuration:", err)
	}

	return cfg
}

func mustEnv(key string) string {
	val := strings.TrimSpace(os.Getenv(key))
	if val == "" {
		log.Fatalf("missing env var: %s", key)
	}
	return val
}

func mustDuration(key string, def time.Duration) time.Duration {
	raw := strings.TrimSpace(os.Getenv(key))
	if raw == "" {
		return def
	}

	d, err := time.ParseDuration(raw)
	if err != nil {
		log.Fatalf("invalid duration for %s: %v", key, err)
	}
	return d
}

func mustInt32(key string, def int32) int32 {
	raw := strings.TrimSpace(os.Getenv(key))
	if raw == "" {
		return def
	}

	v, err := strconv.Atoi(raw)
	if err != nil {
		log.Fatalf("invalid int for %s: %v", key, err)
	}
	return int32(v)
}

func parseCSV(v string) []string {
	parts := strings.Split(v, ",")
	out := make([]string, 0, len(parts))

	for _, p := range parts {
		if s := strings.TrimSpace(p); s != "" {
			out = append(out, s)
		}
	}

	return out
}
