package main

import (
	"log"
	"travel/api"
	"travel/config"
	"travel/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	r := api.NewRouter(db)
	r.Run(":" + config.Load().AUTH_SERVICE_PORT)
}
