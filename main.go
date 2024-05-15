package main

import (
	"context"
	"log"

	"github.com/Emmanuel-MacAnThony/simple_bank/api"
	db "github.com/Emmanuel-MacAnThony/simple_bank/db/sqlc"
	"github.com/Emmanuel-MacAnThony/simple_bank/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot Load Config")
	}
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db")
	}
	log.Println("Connected to DB 🚀🚀🚀")
	store := db.NewStore(connPool)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server")
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server")
	}

}
