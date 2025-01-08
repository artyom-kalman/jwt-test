package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

func AccessToken(rw http.ResponseWriter, r *http.Request) {
	userId, err := getUserIdFromRequest(r)
	if err != nil {
		http.Error(rw, "error reading request body", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	log.Printf("Attempt to get token: %s", userId)

	user, err := getUserById(userId)
	if err != nil {
		http.Error(rw, "invalid user id", http.StatusForbidden)
		log.Printf("Unsuccessful attempt: %s", userId)
		return
	}
	user.Ip = r.RemoteAddr

	accessToken, refreshToken, err := generateTokens(user)
	if err != nil {
		http.Error(rw, "error authenticating user", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	loginResponse := NewLoginResponse(accessToken, refreshToken)

	jsonData, err := json.Marshal(loginResponse)
	if err != nil {
		http.Error(rw, "error authentication user", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jsonData)
}

func getUserIdFromRequest(r *http.Request) (string, error) {
	userId := r.URL.Query().Get("id")
	if userId == "" {
		return "", errors.New("Invalid request: expected id parameter")
	}

	return userId, nil
}

func getUserById(userId string) (*UserData, error) {
	users := DatabaseFabric("database/users.sql")

	userExists, err := users.IsUserIdValid(userId)
	if err != nil {
		return nil, err
	}
	if !userExists {
		return nil, nil
	}

	user, err := users.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func generateTokens(userData *UserData) (string, string, error) {
	iat := time.Now().Unix()
	accessToken, err := CreateAccessToken(userData, iat)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := CreateRefreshToken(userData, iat)
	if err != nil {
		return "", "", err
	}

	db := NewAuthDB("database/users.sql")
	err = db.InsertRefteshToken(userData, refreshToken)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
