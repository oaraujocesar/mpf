package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/oaraujocesar/mpf/configs"
)

func ConnectDB() (*pgx.Conn, context.Context) {
	ctx := context.Background()
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	conn, err := pgx.Connect(ctx, config.DB_URL)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
	}

	return conn, ctx
}
