package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppPort    string
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	DBSSLMode  string
	JWTSecret  string
}

func Load() Config {
	cfg := Config{
		AppPort:    getEnv("APP_PORT", "8080"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBName:     getEnv("DB_NAME", "iot_flutter_app"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBSSLMode:  getEnv("DB_SSL_MODE", "disable"),
		JWTSecret:  mustGetEnv("JWT_SECRET"),
	}

	return cfg
}

func (c Config) DatabaseURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
		c.DBSSLMode,
	)
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}

func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("missing required environment variable: %s", key))
	}

	return value
}
