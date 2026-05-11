package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type DB interface {
	Connect(ctx context.Context, connString string) (*pgx.Conn, error)
	CreateBasicTable(ctx context.Context, conn *pgx.Conn) (err error)
}
