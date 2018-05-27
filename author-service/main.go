package main

import (
	"fmt"
	"log"
	"net"
	"os"

	proto "github.com/tamarakaufler/grpc-publication-manager/author-service/proto"
	"google.golang.org/grpc"
)

var (
	port string = ":50051"
)

func main() {
	log.Println("Starting author-service ...")

	dbConn, err := DBConnection()
	defer dbConn.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	} else {
		log.Println("Connected to database ...")
	}

	// Mmigrates author struct
	// into database columns/types.
	// Will migrate changes each time
	// the service is restarted.
	// dbConn.AutoMigrate(&proto.Author{})

	db := &Store{dbConn}
	tokenService := TokenService{}

	// gRPC server
	customPort := os.Getenv("SERVICE_PORT")
	if customPort != "" {
		port = fmt.Sprintf(":%s", customPort)
	}

	conn, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Register the author service with the gRPC server.
	// Ties our implementation with the auto-generated interface code
	// for the protobuf definition.
	proto.RegisterAuthorServiceServer(s, &service{db, tokenService})

	log.Println("Starting server ...")
	if err := s.Serve(conn); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
