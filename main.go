package main

import (
	"log"

	"github.com/koolman42/go-microservice/database"
	"github.com/koolman42/go-microservice/server"
)

func main() {

	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("failed to initialize Database Client: %s", err)
	}
	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
