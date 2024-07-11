package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

// App config
type Config struct {
	IsDebug        bool `env:"IS_DEBUG" envDefault:"false"`
	IsDevelopment  bool `env:"IS_DEVELOPMENT" envDefault:"false"`
	ServerConfig   HTTPServerConfig
	DBconfig       DBConfig `env-prefix:"DB_"`
	MigrationsPath string
}

type HTTPServerConfig struct {
	Address     string        `env:"ADDRESS" env-default:":localhost"`
	Timeout     time.Duration `env:"TIMEOUT" env-default:"5s"`
	IdleTimeout time.Duration `env:"IDLE_TIMEOUT" env-default:"5s"`
}

type DBConfig struct {
	Host            string        `env:"HOST" envDefault:"localhost"`
	Port            uint          `env:"PORT" envDefault:"5432"`
	Login           string        `env:"LOGIN" envDefault:"postgres"`
	DBName          string        `env:"NAME" envDefault:"postgres"`
	Password        string        `env:"PASSWORD" envDefault:"postgres"`
	SSLMode         string        `env:"SSLMODE" envDefault:"disable"`
	Schema          string        `env:"SCHEMA" envDefault:"public"`
	MaxOpenConns    uint          `env:"MAX_OPEN_CONNS" envDefault:"10"`
	ConnMaxLifeTime time.Duration `env:"CONN_MAX_LIFE_TIME" envDefault:"10s"`
}

// Load - reads config from environment, else fatal
func Load() (*Config, error) {
	cfg := &Config{}

	log.Printf("trying to load config from env")

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}

	if migrationPath, exists := os.LookupEnv("MIGRATIONS_PATH"); exists {
		cfg.MigrationsPath = migrationPath
	}

	return cfg, nil
}
