package auth

import (
	"fmt"

	"github.com/mitch000001/timetables/Godeps/_workspace/src/code.google.com/p/goauth2/oauth"
	"github.com/mitch000001/timetables/Godeps/_workspace/src/github.com/mitch000001/go-harvest/harvest"
	"github.com/mitch000001/timetables/Godeps/_workspace/src/golang.org/x/oauth2"
)

var oauth2AuthURLTemplate string = "%s/oauth2/authorize"
var oauth2TokenURLTemplate string = "%s/oauth2/token"

// TODO(mw): Handling for wrong subdomains
func NewOauth2EndpointForSubdomain(subdomain string) oauth2.Endpoint {
	return oauth2.Endpoint{
		AuthURL:  Oauth2AuthURL(subdomain),
		TokenURL: Oauth2TokenURL(subdomain),
	}
}

func Oauth2AuthURL(subdomain string) string {
	return fmt.Sprintf(oauth2AuthURLTemplate, subdomain)
}

func Oauth2TokenURL(subdomain string) string {
	return fmt.Sprintf(oauth2TokenURLTemplate, subdomain)
}

// NewBasicAuthClient creates a new ClientProvider with BasicAuth as authentication method
func NewBasicAuthClientProvider(config *BasicAuthConfig) harvest.HttpClientProvider {
	clientProvider := &Transport{Config: config}
	return harvest.ClientProviderFunc(clientProvider.Client)
}

// NewOAuthClient creates a new ClientProvider with OAuth as authentication method
func NewOAuthClientProvider(config *oauth.Config) harvest.HttpClientProvider {
	clientProvider := &oauth.Transport{Config: config}
	return harvest.ClientProviderFunc(clientProvider.Client)
}
