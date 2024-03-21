package main

import (
	"cloudplane/api"
	db "cloudplane/db/model"
	"cloudplane/internal"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// @securityDefinitions.apikey	BearerAuth
// @name						Authorization
// @in							header
// @tokenurl					http://127.0.0.1:9090/api/user/login
func main() {
	config, err := internal.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	sqlConn, err := sql.Open(config.DBDriver, config.DBConn)
	if err != nil {
		panic(err)
	}

	store := db.SetupDB(sqlConn)

	server := api.NewServer(config, store)
	log.Println(server.Run())
}
