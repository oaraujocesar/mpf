package database

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://postgres:postgres@database:5432/mpf_test?sslmode=disable"
)

var testQueries *Queries
var migrations *migrate.Migrate

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}

	testQueries = New(conn)
	migrations = setupMigrations(conn)

	os.Exit(m.Run())
}

func setupMigrations(conn *sql.DB) *migrate.Migrate {
	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}

	migrate, err := migrate.NewWithDatabaseInstance("file://../migrations", dbDriver, driver)
	if err != nil {
		log.Fatal("cannot migrate database: ", err)
	}

	return migrate
}

func setupTest(migrate *migrate.Migrate) {
	migrate.Up()
}

func teardownTest(migrate *migrate.Migrate) {
	migrate.Down()
}
