package main

import (
	"io"
	"net/http"
)
func GetToken(wr http.ResponseWriter, r *http.Request) {
	io.WriteString(wr, "Token here")
}
