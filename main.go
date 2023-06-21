package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/wanderlust-create/adventure-life-simple-bank/api"
	db "github.com/wanderlust-create/adventure-life-simple-bank/db/simple_bank_db"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5433/adventure_life_simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
