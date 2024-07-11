package pg

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"movie-project/internal/app/config"
	"time"
)

const (
	maxConnIdleTimeDefault = 5 * time.Second
	MaxOpenConnsModifier   = 2
)

func GetDB(ctx context.Context, cfg config.DBConfig) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s search_path=%s default_query_exec_mode=simple_protocol",
		cfg.Host,
		cfg.Port,
		cfg.Login,
		cfg.DBName,
		cfg.Password,
		cfg.SSLMode,
		cfg.Schema,
	)

	dbconfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parce db connection string: %w", err)
	}

	if cfg.ConnMaxLifeTime > 0 {
		dbconfig.MaxConnLifetime = cfg.ConnMaxLifeTime
	}

	if cfg.MaxOpenConns > 0 {
		dbconfig.MaxConns = int32(cfg.MaxOpenConns / MaxOpenConnsModifier)
	}

	dbconfig.MaxConnIdleTime = maxConnIdleTimeDefault

	pool, err := pgxpool.NewWithConfig(ctx, dbconfig)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize pgx pool: %w", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return pool, nil
}
