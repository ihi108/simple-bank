package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"simple-bank/util"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the database: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
