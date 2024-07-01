package config

import "os"

// Server config
type Config struct {
	HTTPAddr       string
	DSN            string
	MigrationsPath string
}

// Read - reads config from environment
func Read() Config {
	var cfg Config

	httpAddr, exists := os.LookupEnv("HTTP_ADDR")
	if exists {
		cfg.HTTPAddr = httpAddr
	}

	DSN, exists := os.LookupEnv("DSN")
	if exists {
		cfg.DSN = DSN
	}

	migrationsPath, exists := os.LookupEnv("MIGRATIONS_PATH")
	if exists {
		cfg.MigrationsPath = migrationsPath
	}

	return Config{}
}
