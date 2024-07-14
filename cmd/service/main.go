package main

import (
	"fmt"
	"log"
	"net"
	"travel/config"
	pb "travel/genproto/users"
	"travel/service"
	"travel/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", config.Load().USER_SERVICE_PORT)
	if err != nil {
		log.Panic(err)
	}
	defer listener.Close()

	db, err := postgres.ConnectDB()
	if err != nil {
		log.Panic(err)
	}

	u := service.NewUserService(db)
	server := grpc.NewServer()
	pb.RegisterUsersServer(server, u)

	fmt.Printf("User service is listening on port %s...\n", config.Load().USER_SERVICE_PORT)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Error with listening user server: %s", err)
	}
}