package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PGHost   string
	PGPort   string
	PGUser   string
	PGPass   string
	PGDBName string
}

func LoadConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	return &Config{
		PGHost: getEnv("PG_HOST"),
		PGPort: getEnv("PG_PORT"),
		PGUser: getEnv("PG_USERNAME"),
		PGPass: getEnv("PG_PASSWORD"),
		PGDBName: getEnv("PG_DB"),
	}
}

func getEnv(key string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return ""
}