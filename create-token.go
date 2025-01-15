package main

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var secret = []byte("secret")

func CreateAccessToken(userData *UserData, iat int64) (string, string, error) {
	tokenId := uuid.NewString()

	claims := jwt.MapClaims{
		"clientIp": userData.Ip,
		"exp":      time.Now().Add(time.Hour).Unix(),
		"id":       tokenId,
	}

	token, err := createToken(claims)
	if err != nil {
		return "", "", err
	}

	return token, tokenId, nil
}

func CreateRefreshToken(userData *UserData, tokenId string) (string, error) {
	claims := jwt.MapClaims{
		"clientIp": userData.Ip,
		"exp":      time.Now().Add(time.Hour * 48).Unix(),
		"tokenId":  tokenId,
	}

	return createToken(claims)
}

func createToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS512,
		claims,
	)

	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
