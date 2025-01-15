package main

import "encoding/base64"

type IssueResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func NewIssueResponse(accessToken string, refreshToken string) *IssueResponse {
	return &IssueResponse{
		AccessToken:  accessToken,
		RefreshToken: base64.StdEncoding.EncodeToString([]byte(refreshToken)),
	}
}
