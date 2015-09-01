package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"net/url"

	"github.com/mitch000001/go-harvest/harvest"
	"golang.org/x/oauth2"
)

type SessionManager map[string]*Session

func (s *SessionManager) init() {
	if s == nil {
		*s = make(map[string]*Session)
	}
}

func (sm *SessionManager) Add(s *Session) {
	sm.init()
	(*sm)[s.id] = s
}

func (sm *SessionManager) Find(sessionId string) *Session {
	return (*sm)[sessionId]
}

func (sm *SessionManager) Remove(s *Session) {
	delete(*sm, s.id)
}

type Session struct {
	Stack       string
	URL         *url.URL
	location    string
	googleToken *oauth2.Token
	User        *User
	id          string
	errors      []error
}

func (s *Session) LoggedIn() bool {
	return s.User != nil
}

func (s *Session) GetHarvestClient() (*harvest.Harvest, error) {
	config := s.User.HarvestOauth2Config()
	if config == nil {
		return nil, fmt.Errorf("Missing harvest oauth config")
	}
	token := s.User.HarvestToken()
	if token == nil {
		return nil, fmt.Errorf("Missing harvest token")
	}
	// TODO(mw): validate that the token is valid and if not, exchange a new token!
	client, err := harvest.New(config.Subdomain, func() harvest.HttpClient { return config.Client(oauth2.NoContext, token) })
	if err != nil {
		return nil, fmt.Errorf("Error while creating new harvest client: %T(%v)", err, err)
	}
	return client, nil
}

func (s *Session) AddError(err error) {
	if s.errors == nil {
		s.errors = make([]error, 0)
	}
	s.errors = append(s.errors, err)
}

func (s *Session) AddDebugError(err error) {
	if debugMode {
		s.AddError(err)
	}
}

func (s *Session) GetErrors() []error {
	return s.errors
}

func (s *Session) ResetErrors() {
	s.errors = make([]error, 0)
}

func newSession() *Session {
	b := make([]byte, 30)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	id := fmt.Sprintf("%x", sha256.Sum256(b))
	return &Session{id: id}
}
