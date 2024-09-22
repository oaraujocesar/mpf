package main

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/oaraujocesar/mpf/cmd/migrations"
	"github.com/oaraujocesar/mpf/configs"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("pgx", config.DBUrl)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	migrations.Up(db)
}
