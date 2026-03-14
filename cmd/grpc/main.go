package main

import (
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"github.com/aryonuwi/tablink-test-BE/internal/pkg/config"
)

func main() {
	_ = godotenv.Load()

	cfg := config.Load()

	lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	log.Printf("gRPC server running on :%s", cfg.GRPCPort)
	log.Fatal(server.Serve(lis))
}
