package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/wanderlust-create/adventure-life-simple-bank/api"
	db "github.com/wanderlust-create/adventure-life-simple-bank/db/simple_bank_db"
	"github.com/wanderlust-create/adventure-life-simple-bank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)// connects to the real SQL db
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
