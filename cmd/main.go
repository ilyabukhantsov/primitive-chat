package main

import (
	"context"
	"log"
	"primitive-chat/internal/db"
)

func main() {

	ctx := context.Background()

	// Connection string: postgres://user:password@host:port/dbname
	connString := "postgres://user:password@db:5432/mydb"

	dbService := db.NewDBService(nil)
	conn, err := dbService.Connect(ctx, connString)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close(ctx)

	dbService.CreateBasicTable(ctx, conn)
}
