package main

import (
	"log"
	"net/http"
)

func main() {
	const PORT = ":3030"

	http.HandleFunc("/access-token", IssueTokenHandler)
	http.HandleFunc("/refresh-token", RefreshTokenHandler)

	log.Printf("Runnign server on %s", PORT)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		panic(err)
	}
}
