package main

import (
	"io"
	"net/http"
)

func RefreshToken(wr http.ResponseWriter, r *http.Request) {
	io.WriteString(wr, "New token")
}
