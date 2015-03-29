package main

import (
	"github.com/mitch000001/timetables/Godeps/_workspace/src/github.com/mitch000001/go-harvest/harvest"
	"github.com/mitch000001/timetables/Godeps/_workspace/src/golang.org/x/oauth2"
)

func NewUser(idToken *googleIdToken) *User {
	var company *Company
	if idToken.HostedDomain != "" {
		company = &Company{Domain: idToken.HostedDomain}
	}
	return &User{idToken: idToken, company: company}
}

type User struct {
	idToken *googleIdToken
	profile *googleProfile
	company *Company
	*harvest.AccountUser
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

func (u *User) SetHarvestAccount(account *harvest.Account) {
	if u.company != nil {
		u.company.Account = account
	}
	u.AccountUser = account.User
}

type Company struct {
	*harvest.Account
	Domain              string
	HarvestSubdomain    string
	harvestOauth2Config *oauth2.Config
	harvestToken        *oauth2.Token
}
