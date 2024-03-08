package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewConnection() *Repository {
	dbUrl := os.Getenv("DB_CONN_STR")
	config, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse config: %v\n", err)
		os.Exit(1)
	}

	config.MinConns = 5
	config.MaxConns = 10

	pool, err := pgxpool.NewWithConfig(
		context.Background(),
		config,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return &Repository{
		pool: pool,
	}
}

func (repo *Repository) CloseConnection() {
	repo.pool.Close()
}
