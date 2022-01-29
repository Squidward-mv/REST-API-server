package sql_store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost dbname=restapi_test user=slaine password=wasp1488 sslmode=disable"
	}

	os.Exit(m.Run())
}
