package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/LeMinh0706/simplebank/api"
	db "github.com/LeMinh0706/simplebank/db/sqlc"
	"github.com/LeMinh0706/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Can't connect to DB:", err)
	}

	queries := db.New(conn)
	server, err := api.NewServer(config, queries)
	if err != nil {
		log.Fatal("Cannot create server:", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Can't start server:", err)
	}
}
