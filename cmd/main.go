package main

import (
	"context"
	pb "github.com/pramonow/go-grpc-server-streaming-example/src/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"primitive-chat/internal/db"
	"primitive-chat/internal/grpc_server"
)

func main() {

	ctx := context.Background()

	// Connection string: postgres://user:password@host:port/dbname
	connString := "postgres://user:password@db:5432/mydb"

	server := grpc_server.NewServerStruct()
	dbService := db.NewDBService(nil)
	conn, err := dbService.Connect(ctx, connString)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = conn.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = dbService.CreateBasicTable(ctx, conn)
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", ":50005")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create grpc server
	s := grpc.NewServer()
	pb.RegisterStreamServiceServer(s, server)

	log.Println("start server")
	// and start...
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
