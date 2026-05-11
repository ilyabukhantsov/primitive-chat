package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type Postgress struct {
	db DB
}

func (p *Postgress) Connect(ctx context.Context, connString string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully connected!")

	return conn, nil
}

func (p *Postgress) CreateBasicTable(ctx context.Context, conn *pgx.Conn) (err error) {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`
	_, err = conn.Exec(ctx, createTableSQL)
	if err != nil {
		return err
	}
	fmt.Println("Table created successfully!")
	return nil
}

func NewDBService(d DB) *Postgress {
	return &Postgress{
		db: d,
	}
}
