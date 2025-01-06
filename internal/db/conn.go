package db

import (
	"time"

	"github.com/arafetki/go-echo-boilerplate/internal/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	*pgxpool.Pool
}

func Pool(dsn string) (*DB, error) {

	ctx, cancel := utils.ContextWithTimeout(3 * time.Second)
	defer cancel()

	config, err := pgxpool.ParseConfig("postgresql://" + dsn)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	return &DB{pool}, nil
}
