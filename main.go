package main

import (
	"net/http"
)

func main() {
	const PORT = ":3030"

	http.HandleFunc("/access-token", AccessToken)
	http.HandleFunc("/refresh-token", RefreshToken)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		panic(err)
	}
}
