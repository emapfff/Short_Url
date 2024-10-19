package config

import (
	"os"
)

type DatabaseConfig struct {
	DbUser, DbPassword, DbName, DbHost, DbPort string
}

type Config struct {
	Database DatabaseConfig
}

func Init() *Config {
	return &Config{
		Database: DatabaseConfig{
			DbUser:     getEnv("POSTGRES_USER", ""),
			DbPassword: getEnv("POSTGRES_PASSWORD", ""),
			DbName:     getEnv("POSTGRES_DB", ""),
			DbHost:     getEnv("POSTGRES_HOST", ""),
			DbPort:     getEnv("POSTGRES_PORT", ""),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
