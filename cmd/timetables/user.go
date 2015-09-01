package main

import (
	"net/http"

	"github.com/mitch000001/go-harvest/harvest"
	"golang.org/x/oauth2"
)

func NewUser(idToken *googleIdToken) *User {
	company := CompanyForGoogleToken(idToken)
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

func (u *User) FullName() string {
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

func (u *User) SetHarvestOauthConfig(config *HarvestOauth2Config) {
	u.company.harvestOauth2Config = config
}

func (u *User) HarvestOauth2Config() *HarvestOauth2Config {
	return u.company.harvestOauth2Config
}

func (u *User) SetHarvestToken(token *oauth2.Token) {
	u.company.harvestToken = token
}

func (u *User) HarvestToken() *oauth2.Token {
	return u.company.harvestToken
}

func (u *User) HarvestSubdomain() string {
	return u.company.harvestOauth2Config.Subdomain
}

func (u *User) SetHarvestAccount(account *harvest.Account) {
	if u.company != nil {
		u.company.Account = account
	}
	u.AccountUser = account.User
}

func usersShowHandler() authHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, s *Session) {
		page := PageForSession(s)
		page.Set("User", s.User)
		renderTemplate(w, "users-show", page)
	}
}

var CompanyRepository Companies = make(Companies)

type Companies map[string]*Company

func (c *Companies) FindByDomain(domain string) (*Company, bool) {
	if domain == "" {
		return nil, false
	}
	company, ok := (*c)[domain]
	return company, ok
}

func (c *Companies) Add(company *Company) bool {
	_, ok := (*c)[company.Domain]
	if ok {
		return false
	}
	(*c)[company.Domain] = company
	return true
}

func CompanyForGoogleToken(idToken *googleIdToken) *Company {
	company, ok := CompanyRepository.FindByDomain(idToken.HostedDomain)
	if ok {
		return company
	}
	if idToken.HostedDomain != "" {
		company = &Company{Domain: idToken.HostedDomain}
	} else {
		company = &Company{Domain: idToken.Email}
	}
	CompanyRepository.Add(company)
	return company
}

type Company struct {
	*harvest.Account
	Domain              string
	harvestOauth2Config *HarvestOauth2Config
	harvestToken        *oauth2.Token
}
