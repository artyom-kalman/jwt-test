package main

import "github.com/golang-jwt/jwt/v5"

var secret = []byte("secret")

func createToken(userData *UserData) (string, error) {
	claims := jwt.MapClaims{
		"userIp":   userData.Ip,
		"username": userData.Username,
	}

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
