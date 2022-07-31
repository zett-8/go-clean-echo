package configs

import (
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"time"
)

type Auth0ConfigType struct {
	Domain             string
	ClientID           string
	Audience           []string
	Issuer             string
	SignatureAlgorithm validator.SignatureAlgorithm
	CacheDuration      time.Duration
}

var Auth0Config = Auth0ConfigType{
	Domain:             "****",
	ClientID:           "****",
	Audience:           []string{"****"},
	Issuer:             "****",
	SignatureAlgorithm: validator.RS256,
	CacheDuration:      15 * time.Minute,
}
