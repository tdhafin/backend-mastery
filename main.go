package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/tdhafin/backend-mastery/api"
	db "github.com/tdhafin/backend-mastery/db/sqlc"
	"github.com/tdhafin/backend-mastery/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("caanot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database")
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
