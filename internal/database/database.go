package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Connection struct {
    Context context.Context
    Connection *pgxpool.Pool
}

func Connect(ctx context.Context, databaseURL string) (*Connection, error) {
    conn, err := pgxpool.New(ctx, databaseURL)

    if err != nil {
        return nil, err
    }

    return &Connection{
        Context: ctx,
        Connection: conn,
    }, nil
}
