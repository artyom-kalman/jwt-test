package main

import "encoding/base64"

type IssueResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewIssueResponse(accessToken string, refreshToken string) *IssueResponse {
	return &IssueResponse{
		AccessToken:  accessToken,
		RefreshToken: base64.StdEncoding.EncodeToString([]byte(refreshToken)),
	}
}
