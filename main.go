package main

import (
	"context"
	"log"

	"github.com/Emmanuel-MacAnThony/simple_bank/api"
	db "github.com/Emmanuel-MacAnThony/simple_bank/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:feb061999@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	connPool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db")
	}
	log.Println("Connected to DB 🚀🚀🚀")
	store := db.NewStore(connPool)
	server, err := api.NewServer(store)
	if err != nil {
		log.Fatal("Cannot create server")
	}
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start server")
	}

}
