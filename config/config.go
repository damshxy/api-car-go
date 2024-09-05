package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PGHOST string
	PGPORT string
	PGUSER string
	PGPASSWORD string
	PGDATABASE string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	return &Config{
		PGHOST: getEnv("PG_HOST"),
		PGPORT: getEnv("PG_PORT"),
		PGUSER: getEnv("PG_USERNAME"),
		PGPASSWORD: getEnv("PG_PASSWORD"),
		PGDATABASE: getEnv("PG_DATABASE"),
	}
}

func getEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}

	return val
}