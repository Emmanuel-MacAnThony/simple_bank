package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/Emmanuel-MacAnThony/simple_bank/utils"
	"github.com/jackc/pgx/v5/pgxpool"
	// _ "github.com/lib/pq"
)

var testStore Store
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot Load Config")
	}
	ctx := context.Background()
	testDB, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)

	}
	defer testDB.Close()
	testStore = NewStore(testDB)

	os.Exit(m.Run())

}
