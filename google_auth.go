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
		r.ParseForm()
		form := r.PostForm
		s := newSession()
		s.location = form.Get("referer")
		debug.Printf("form: %+#v\n", form)
		debug.Printf("Referer: %s\n", s.location)
		sessions.Add(s)
		url := config.AuthCodeURL(s.id, oauth2.AccessTypeOffline)
		copyHeader(w.Header(), r.Header)
		http.Redirect(w, r, url, http.StatusFound)
		return
	}
}

func googleRedirectHandler(config *oauth2.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		state := params.Get("state")
		if state == "" {
			copyHeader(w.Header(), r.Header)
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		session := sessions.Find(state)
		if session == nil {
			copyHeader(w.Header(), r.Header)
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		code := params.Get("code")
		if code == "" {
			copyHeader(w.Header(), r.Header)
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		token, err := config.Exchange(oauth2.NoContext, code)
		if err != nil {
			copyHeader(w.Header(), r.Header)
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		session.googleToken = token
		id := token.Extra("id_token")
		idToken, err := decode(id.(string))
		if err != nil {
			copyHeader(w.Header(), r.Header)
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		session.idToken = idToken
		http.SetCookie(w, newSessionCookie(session.id))
		copyHeader(w.Header(), r.Header)
		http.Redirect(w, r, session.location, http.StatusFound)
	}
}

func newSessionCookie(sessionId string) *http.Cookie {
	return &http.Cookie{Name: "timetable", Value: sessionId, Expires: time.Now().Add(5 * 24 * time.Hour)}
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
