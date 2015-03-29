package main

import (
	"github.com/mitch000001/timetables/Godeps/_workspace/src/github.com/mitch000001/go-harvest/harvest"
	"golang.org/x/oauth2"
)

func NewUser(idToken *googleIdToken) *User {
	return &User{idToken: idToken}
}

type User struct {
	idToken *googleIdToken
	profile *googleProfile
	*harvest.User
	backOffice bool
	admin      bool
}

func (u *User) SetProfile(profile *googleProfile) {
	u.profile = profile
}

func (u *User) Email() string {
	return u.idToken.Email
}

func (u *User) String() string {
	if u.profile != nil {
		return u.profile.FullName()
	}
	return u.idToken.Email
}

func (u *User) IsBackOffice() bool {
	return u.backOffice
}

func (u *User) IsAdmin() bool {
	return u.admin
}

type Company struct {
	Domain              string
	HarvestSubdomain    string
	harvestOauth2Config *oauth2.Config
	harvestToken        *oauth2.Token
}
