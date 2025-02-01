package main

type UserData struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	ClientIp string
}
