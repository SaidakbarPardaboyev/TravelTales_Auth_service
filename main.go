package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"travel/api"
	"travel/config"
	pb "travel/genproto/users"
	"travel/service"
	"travel/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go ServiceRun(&wg)
	go RouterRun(&wg)

	wg.Wait()
}

func ServiceRun(wg *sync.WaitGroup) {
	defer wg.Done()
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

func RouterRun(wg *sync.WaitGroup) {
	defer wg.Done()
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	r := api.NewRouter(db)
	r.Run(config.Load().AUTH_SERVICE_PORT)
}
