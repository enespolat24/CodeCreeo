package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DbConnection struct {
	pool *pgxpool.Pool
}

func NewDbConnection() *DbConnection {
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

	return &DbConnection{
		pool: pool,
	}

}

func (repo *DbConnection) CloseDbConnection() {
	repo.pool.Close()
}
