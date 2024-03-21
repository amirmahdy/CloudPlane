package db

import (
	"cloudplane/internal"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testStore Store

func TestMain(m *testing.M) {
	cfg, err := internal.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load env", err)
	}
	conn, err := sql.Open(cfg.DBDriver, cfg.DBConn)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	defer conn.Close()

	testStore = SetupDB(conn)
	os.Exit(m.Run())
}
