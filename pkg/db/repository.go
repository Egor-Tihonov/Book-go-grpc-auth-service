package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DBPostgres struct {
	Pool *pgxpool.Pool
}

func New(conn string) (*DBPostgres, error) {
	pool, err := pgxpool.Connect(context.Background(), conn)
	if err != nil {
		return nil, err
	}

	return &DBPostgres{Pool: pool}, nil
}
