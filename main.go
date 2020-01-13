package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

// GAEWebURI is GAE target URI
// ClientID is GAE IAP ClientID
// SeviceAccountKeyPath is service account key path
var (
	GAEWebURI            = os.Getenv("GAE_WEB_URI")
	ClientID             = os.Getenv("CLIENT_ID")
	SeviceAccountKeyPath = os.Getenv("SEVICE_ACCOUNT_KEY_PATH")
)

func main() {
	data, _ := ioutil.ReadFile(SeviceAccountKeyPath)

	cfg, _ := google.JWTConfigFromJSON(data, "")
	jwtConfig := &jwt.Config{
		Email:         cfg.Email,
		PrivateKey:    cfg.PrivateKey,
		TokenURL:      cfg.TokenURL,
		Subject:       cfg.Email,
		PrivateClaims: map[string]interface{}{"target_audience": ClientID},
	}

	tokenSource := jwtConfig.TokenSource(oauth2.NoContext)
	token, _ := tokenSource.Token()
	idToken := fmt.Sprintf("%s", token.Extra("id_token"))

	client := &http.Client{}
	req, _ := http.NewRequest("GET", GAEWebURI, nil)
	req.Header.Add("Authorization", "Bearer "+idToken)
	resp, err := client.Do(req)

	if err != nil {
		log.Printf("Error: %s", err)
		return
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	log.Printf("status is %d", resp.StatusCode)
	log.Printf("body is %s", bodyBytes)
}
