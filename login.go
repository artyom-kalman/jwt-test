package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func Login(rw http.ResponseWriter, r *http.Request) {
	userData, err := getUserDataFromRequest(r)
	if err != nil {
		http.Error(rw, "error reading body", http.StatusInternalServerError)
		log.Fatal(err)
	}
	log.Printf("Login attempt: %s, %s", userData.Username, userData.Password)

	isLogedIn, err := authanticateUser(userData)
	if err != nil {
		http.Error(rw, "error authentication user", http.StatusInternalServerError)
		log.Fatal(err)
	}
	if !isLogedIn {
		http.Error(rw, "wrong username of password", http.StatusForbidden)
		log.Printf("Unsuccessful login: %s", userData.Username)
		return
	}
	log.Printf("Successful login: %s", userData.Username)

	accessToken, err := createToken(userData)
	if err != nil {
		http.Error(rw, "error authenticating user", http.StatusInternalServerError)
		log.Fatal(err)
	}

	loginResponse := NewLoginResponse(accessToken, "")

	jsonData, err := json.Marshal(loginResponse)
	if err != nil {
		http.Error(rw, "error authentication user", http.StatusInternalServerError)
		log.Fatalln(err)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jsonData)
}

func getUserDataFromRequest(r *http.Request) (*UserData, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var userData *UserData
	err = json.Unmarshal(body, &userData)
	if err != nil {
		return nil, err
	}

	return userData, nil
}

func authanticateUser(userData *UserData) (bool, error) {
	users := DatabaseFabric("database/users.sql")

	userExists, err := users.UserExists(userData)
	if err != nil {
		return false, err
	}
	if !userExists {
		return false, nil
	}

	isCorrectPassword, err := users.IsPasswordCorrect(userData)
	if err != nil {
		return false, err
	}
	if !isCorrectPassword {
		return false, nil
	}

	return true, nil
}
