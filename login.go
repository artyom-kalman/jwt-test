package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func Login(wr http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(wr, "error reading body", http.StatusInternalServerError)
		log.Fatalln(err)
		return
	}

	var userData *UserData
	err = json.Unmarshal(body, &userData)
	if err != nil {
		http.Error(wr, "error reading body", http.StatusInternalServerError)
		log.Fatalln(err)
		return
	}

	log.Printf("Login attempt: %s, %s", userData.Username, userData.Password)

	users := DatabaseFabric("database/users.sql")

	userExists, err := users.UserExists(userData)
	if err != nil {
		http.Error(wr, "error checking username", http.StatusInternalServerError)
		log.Fatalln(err)
		return
	}
	if !userExists {
		http.Error(wr, "wrong username", http.StatusForbidden)
		return
	}

	isCorrectPassword, err := users.IsPasswordCorrect(userData)
	if err != nil {
		http.Error(wr, "error checking password", http.StatusInternalServerError)
		log.Fatalln(err)
		return
	}
	if !isCorrectPassword {
		http.Error(wr, "wrong password", http.StatusForbidden)
		return
	}
	log.Printf("Successful login: %s", userData.Username)

	io.WriteString(wr, "Token here")
}
