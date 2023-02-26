package middlewares

import (
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/labstack/echo/v4"
	"github.com/zett-8/go-clean-echo/configs"
	"github.com/zett-8/go-clean-echo/logger"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"strings"
)

func JwtMiddleware() (echo.MiddlewareFunc, error) {
	auth0Config := configs.Auth0Config

	issuerURL, err := url.Parse(auth0Config.Issuer)
	if err != nil {
		return nil, err
	}

	provider := jwks.NewCachingProvider(issuerURL, auth0Config.CacheDuration)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		auth0Config.SignatureAlgorithm,
		issuerURL.String(),
		auth0Config.Audience,
	)
	if err != nil {
		return nil, err
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			authorization := c.Request().Header.Get("Authorization")
			if authorization == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "No Authorization Header")
			}

			if !strings.HasPrefix(authorization, "Bearer ") {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization Header")
			}

			token := strings.TrimPrefix(authorization, "Bearer ")

			claims, err := jwtValidator.ValidateToken(c.Request().Context(), token)
			if err != nil {
				logger.Error("Invalid Token: ", zap.Error(err))
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Token")
			}

			c.Set("claims", claims.(*validator.ValidatedClaims))

			return next(c)
		}
	}, nil
}
