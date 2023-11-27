package main

import (
	f "fmt"
	"net/http"

	"github.com/oaraujocesar/mpf/configs"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	f.Printf("Server is running on http://localhost:%s...", config.WebServerPort)
	http.ListenAndServe(f.Sprintf(":%s", config.WebServerPort), nil)
}
