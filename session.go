package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"net/url"

	"github.com/mitch000001/timetables/Godeps/_workspace/src/github.com/mitch000001/go-harvest/harvest"
	"github.com/mitch000001/timetables/Godeps/_workspace/src/golang.org/x/oauth2"
)

type session struct {
	Stack    string
	URL      *url.URL
	location string
	User     *User
	id       string
	errors   []error
}

func (s *session) LoggedIn() bool {
	return s.User != nil
}

func (s *session) GetHarvestClient() (*harvest.Harvest, error) {
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

func (s *session) AddError(err error) {
	if s.errors == nil {
		s.errors = make([]error, 0)
	}
	s.errors = append(s.errors, err)
}

func (s *session) AddDebugError(err error) {
	if debugMode {
		s.AddError(err)
	}
}

func (s *session) GetErrors() []error {
	return s.errors
}

func (s *session) ResetErrors() {
	s.errors = make([]error, 0)
}

func newSession() *session {
	b := make([]byte, 30)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	id := fmt.Sprintf("%x", sha256.Sum256(b))
	return &session{id: id}
}
