package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/oaraujocesar/mpf/app"
	"github.com/oaraujocesar/mpf/configs"
	"github.com/oaraujocesar/mpf/database"
	sqlc "github.com/oaraujocesar/mpf/database/sqlc"
)

func main() {

	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	log.Printf("config file successfully loaded as: %v", config)

	conn, ctx := database.ConnectDB()

	defer conn.Close(ctx)

	sqlc.New(conn)

	server := &http.Server{
		Addr:    ":" + config.WebServerPort,
		Handler: app.Router(),
	}

	fmt.Println("Server running on port", config.WebServerPort)
	log.Fatal(server.ListenAndServe())
}
