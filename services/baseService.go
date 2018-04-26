package services

import (
	"context"
	"net/http"

	"github.com/RideShare-Server/log"
	"github.com/RideShare-Server/utils"
	"github.com/juju/errors"
	"github.com/meshhq/funnel"

	"net/url"

	"github.com/meshhq/gohttp"
	"golang.org/x/oauth2/clientcredentials"
)

// Auth represents the data from an OAuth strategy
type Auth struct {
	ClientID       string
	ClientSecret   string
	OAuthToken     string
	TokenURL       string
	Scopes         []string
	EndpointParams url.Values
}

// BaseService defines the base structure of a service.
type BaseService struct {
	Auth
	Name        string
	APIEndpoint string
	Version     string
	Client      *gohttp.Client
	RateLimit   funnel.RateLimitInfo
}

// Execute :
func (s *BaseService) Execute(request *gohttp.Request) (*gohttp.Response, error) {
	response, err := s.Client.Execute(request)
	if err != nil {
		return nil, errors.Trace(err)
	}

	if response.Code != http.StatusOK {
		return nil, errors.Errorf(utils.EncounteredUnprocessableResponseCode, response.Code)
	}
	return response, nil
}

// GetOAuth2Token handles OAuth2.0 authentication and token refresh.
func (s *BaseService) GetOAuth2Token() string {
	ctx := context.Background()
	conf := &clientcredentials.Config{
		ClientID:       s.ClientID,
		ClientSecret:   s.ClientSecret,
		TokenURL:       s.TokenURL,
		Scopes:         s.Scopes,
		EndpointParams: s.EndpointParams,
	}

	token, err := conf.Token(ctx)
	if err != nil {
		log.Error(log.BaseServiceTopic, err)
	}

	return token.AccessToken
}
