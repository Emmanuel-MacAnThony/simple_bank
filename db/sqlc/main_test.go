package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	// _ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:feb061999@localhost:5432/simple_bank?sslmode=disable"
)

var testStore *SQLStore
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	ctx := context.Background()
	testDB, err := pgxpool.New(ctx, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)

	}
	defer testDB.Close()
	testStore = NewStore(testDB)

	os.Exit(m.Run())

}
