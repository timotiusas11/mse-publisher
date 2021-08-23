package infrastructure

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPsqlDB() (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.Connect(context.Background(), "postgres://postgres:admin@localhost:5432/minisearchengine")
	if err != nil {
		return nil, err
	}

	if err = dbpool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return dbpool, nil
}
