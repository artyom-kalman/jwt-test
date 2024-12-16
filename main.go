package main

import (
	"net/http"
)

func main() {
	const PORT = ":3030"

	http.HandleFunc("/getToken", GetToken)
	http.HandleFunc("/refreshToken", RefreshToken)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		panic(err)
	}
}
