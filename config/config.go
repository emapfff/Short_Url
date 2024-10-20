package config

import (
	"fmt"
	"log"
	"os"
)

type StorageMethod int

const (
	DB StorageMethod = iota
	Map
)

func (sm StorageMethod) String() string {
	return [...]string{"DB", "Map"}[sm]
}

func ParseStorageMethod(s string) (StorageMethod, error) {
	switch s {
	case "DB":
		return DB, nil
	case "Map":
		return Map, nil
	default:
		return -1, fmt.Errorf("invalid storage method: %s", s)
	}
}

type DatabaseConfig struct {
	DbUser, DbPassword, DbName, DbHost, DbPort string
}

type ServerConfig struct {
	Port   string
	Domain string
}

type Config struct {
	Database      DatabaseConfig
	Server        ServerConfig
	StorageMethod StorageMethod
}

func Init() *Config {
	str := getEnv("STORAGE_METHOD", "")
	storageMethod, err := ParseStorageMethod(str)
	if err != nil {
		log.Fatal(err)
	}
	return &Config{
		Database: DatabaseConfig{
			DbUser:     getEnv("POSTGRES_USER", ""),
			DbPassword: getEnv("POSTGRES_PASSWORD", ""),
			DbName:     getEnv("POSTGRES_DB", ""),
			DbHost:     getEnv("POSTGRES_HOST", ""),
			DbPort:     getEnv("POSTGRES_PORT", ""),
		},
		Server: ServerConfig{
			Port:   getEnv("SERVER_PORT", ""),
			Domain: getEnv("SERVER_DOMAIN", ""),
		},
		StorageMethod: storageMethod,
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
