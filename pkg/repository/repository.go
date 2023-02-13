package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DBPostgres struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context, conn string) (*DBPostgres, error) {
	pool, err := pgxpool.Connect(ctx, conn)
	if err != nil {
		return nil, err
	}

	return &DBPostgres{Pool: pool}, nil
}
