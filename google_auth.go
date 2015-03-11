package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/jws"
)

type googleIdToken struct {
	jws.ClaimSet
	AccessTokenHash     string `json:"at_hash"`
	EmailVerified       bool   `json:"email_verified"`
	AuthorizedPresenter string `json:"azp"`
	Email               string `json:"email"`
}

func googleLoginHandler(config *oauth2.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s := newSession()
		s.location = r.Header.Get("X-Referer")
		sessions.Add(s)
		url := config.AuthCodeURL(s.id, oauth2.AccessTypeOffline)
		http.Redirect(w, r, url, http.StatusFound)
		return
	}
}

func googleRedirectHandler(config *oauth2.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		state := params.Get("state")
		if state == "" {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		session := sessions.Find(state)
		if session == nil {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		code := params.Get("code")
		if code == "" {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		token, err := config.Exchange(oauth2.NoContext, code)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		session.googleToken = token
		id := token.Extra("id_token")
		idToken, err := decode(id.(string))
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		session.idToken = idToken
		http.SetCookie(w, &http.Cookie{Name: "timetable", Value: session.id, Expires: time.Now().Add(5 * 24 * time.Hour)})
		http.Redirect(w, r, session.location, http.StatusFound)
	}
}

func decode(payload string) (*googleIdToken, error) {
	// decode returned id token to get expiry
	s := strings.Split(payload, ".")
	if len(s) < 2 {
		return nil, errors.New("jws: invalid token received")
	}
	decoded, err := base64Decode(s[1])
	if err != nil {
		return nil, err
	}
	c := &googleIdToken{}
	err = json.NewDecoder(bytes.NewBuffer(decoded)).Decode(c)
	return c, err
}
func base64Decode(s string) ([]byte, error) {
	// add back missing padding
	switch len(s) % 4 {
	case 2:
		s += "=="
	case 3:
		s += "="
	}
	return base64.URLEncoding.DecodeString(s)
}
