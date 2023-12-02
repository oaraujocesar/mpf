package controllers

import (
	"encoding/json"
	"net/http"

	database "github.com/oaraujocesar/mpf/database/sqlc"
)

func Signup(w http.ResponseWriter, r *http.Request) {

	var user database.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// create user

	w.Write([]byte("Hello World!"))
}
