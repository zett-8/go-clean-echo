package configs

import (
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"time"
)

type auth0ConfigSample struct {
	Domain   string
	ClientID string
	// https://sample.us.auth0.com/api/v2/ auth0 menu -> Applications -> APIs -> API Audience
	Audience []string
	// https://<your tenant domain>/
	Issuer             string
	SignatureAlgorithm validator.SignatureAlgorithm
	CacheDuration      time.Duration
}

var Auth0ConfigSample = auth0ConfigSample{
	Domain:             "",
	ClientID:           "",
	Audience:           []string{""},
	Issuer:             "",
	SignatureAlgorithm: validator.RS256,
	CacheDuration:      15 * time.Minute,
}
