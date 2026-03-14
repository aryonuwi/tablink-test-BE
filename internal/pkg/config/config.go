package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppName  string
	AppEnv   string
	HTTPPort string
	GRPCPort string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

func Load() Config {
	return Config{
		AppName:    getEnv("APP_NAME", "tablelink-test"),
		AppEnv:     getEnv("APP_ENV", "local"),
		HTTPPort:   getEnv("HTTP_PORT", "8080"),
		GRPCPort:   getEnv("GRPC_PORT", "50051"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "tablelink_test"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
	}
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
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
