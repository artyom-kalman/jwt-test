package main

import (
	"io"
	"net/http"
)

func RefreshToken(wr http.ResponseWriter, r *http.Request) {
	// get refresh

	// validate refresh

	// generate new pair
	io.WriteString(wr, "New token")
}
