package main

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("secret")

func CreateAccessToken(userData *UserData, iat int64) (string, error) {
	claims := jwt.MapClaims{
		"userId": userData.Id,
		"userIp": userData.Ip,
		"iat":    iat,
		"exp":    time.Now().Add(time.Hour).Unix(),
	}

	return createToken(claims)
}

func CreateRefreshToken(userData *UserData, iat int64) (string, error) {
	claims := jwt.MapClaims{
		"userId": userData.Id,
		"userIp": userData.Ip,
		"iat":    iat,
		"exp":    time.Now().Add(time.Hour * 48).Unix(),
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
