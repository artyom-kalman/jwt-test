package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type RefereshRequestBody struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func RefreshTokenHandler(wr http.ResponseWriter, r *http.Request) {
	refreshBody, err := GetRefreshRequestBody(r)
	if err != nil {
		http.Error(wr, "error reading request body", http.StatusBadRequest)
		log.Printf("error reading refresh request body: %v", err)
		return
	}

	clientIp := getClientIpAddress(r)
	isValid, err := ValidateRefreshToken(refreshBody.AccessToken, refreshBody.RefreshToken, clientIp)
	if err != nil {
		http.Error(wr, "error validating token", http.StatusInternalServerError)
		log.Printf("error validating token: %s", err)
		return
	}
	if !isValid {
		http.Error(wr, "invalid token", http.StatusUnauthorized)
		log.Printf("invalid token: %v", err)
		return
	}

	// generate new pair
	io.WriteString(wr, "New token")
}

func GetRefreshRequestBody(r *http.Request) (*RefereshRequestBody, error) {
	var refreshBody RefereshRequestBody

	err := json.NewDecoder(r.Body).Decode(&refreshBody)
	if err != nil {
		return nil, err
	}

	decodedRefreshToken, err := base64.StdEncoding.DecodeString(refreshBody.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("error decoding refresh token: %v", err)
	}

	refreshBody.RefreshToken = string(decodedRefreshToken)

	return &refreshBody, nil
}
func ValidateRefreshToken(accessToken string, refreshToken string, clientIp string) (bool, error) {
	parsedToken, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return false, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return false, nil
	}

	isClientIdValid := validateClientIp(clientIp, claims)
	if !isClientIdValid {
		return false, nil
	}

	tokenId, ok := claims["token_id"].(string)
	if !ok {
		return false, nil
	}
	fmt.Printf("tokenId: %v", tokenId)

	getRefreshTokenFromDb(tokenId)

	return true, nil
}

func validateClientIp(clientIp string, claims jwt.MapClaims) bool {
	originClientIp, ok := claims["client_id"].(string)
	if !ok {
		return false
	}

	if clientIp != originClientIp {
		return false
	}

	return true
}

func getRefreshTokenFromDb(tokenId string) error {
	db := NewAuthDB("database/users.sql")

	refreshToken, err := db.GetRefreshToken(tokenId)
	if err != nil {
		return err
	}
	if refreshToken == "" {
		return errors.New("invalid token")
	}

	fmt.Printf("refresh token: %s", refreshToken)

	return nil
}
