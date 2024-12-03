package config

import "os"

// Config holds the application configuration
type Config struct {
	PostgresDSN string
	RedisAddr   string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() Config {
	return Config{
		PostgresDSN: os.Getenv("POSTGRES_DSN"),
		RedisAddr:   os.Getenv("REDIS_ADDR"),
	}
}
