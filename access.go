package main

import (
	"encoding/json"
	"errors"
	"log"
	"net"
	"net/http"
	"time"
)

func IssueTokenHandler(rw http.ResponseWriter, r *http.Request) {
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

	user.ClientIp = getClientIpAddress(r)

	accessToken, refreshToken, err := generateTokens(user)
	if err != nil {
		http.Error(rw, "error authenticating user", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	log.Printf("Tokens issued for %s\n", userId)

	issueResponse := NewIssueResponse(accessToken, refreshToken)

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(issueResponse)
}

func getClientIpAddress(r *http.Request) string {
	ipAddress := r.Header.Get("X-Real-Ip")
	if ipAddress == "" {
		ipAddress = r.Header.Get("X-Forwarded-For")
	}
	if ipAddress == "" {
		ipAddress, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	return ipAddress
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
	accessToken, tokenId, err := CreateAccessToken(userData, iat)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := CreateRefreshToken(userData, tokenId)
	if err != nil {
		return "", "", err
	}

	saveRefreshToken(tokenId, userData.Id, refreshToken)

	return accessToken, refreshToken, nil
}

func saveRefreshToken(tokenId string, userId string, token string) error {
	db := NewAuthDB("database/users.sql")

	return db.InsertRefteshToken(tokenId, userId, token)
}
