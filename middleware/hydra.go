package lib

import (
	"fmt"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	hydra "github.com/ory/hydra-client-go"
	"github.com/rs/zerolog/log"
)

type HydraConfig struct {
	Skipper    middleware.Skipper
	Url        string
	Pat        string
	ContextKey string
}

var DefaultHydraConfig = HydraConfig{
	Skipper:    middleware.DefaultSkipper,
	ContextKey: "user",
}

const bearer = "Bearer"

func Hydra() echo.MiddlewareFunc {
	return HydraWithConfig(DefaultHydraConfig)
}

func HydraWithConfig(config HydraConfig) echo.MiddlewareFunc {

	if config.Skipper == nil {
		config.Skipper = DefaultHydraConfig.Skipper
	}

	if config.Url == "" {
		config.Url = os.Getenv("ORY_URL")
	}

	if config.Pat == "" {
		config.Pat = os.Getenv("ORY_PAT")
	}

	conf := hydra.NewConfiguration()
	conf.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %v", config.Pat))
	conf.Servers = []hydra.ServerConfiguration{
		{
			URL: config.Url,
		},
	}
	cli := hydra.NewAPIClient(conf)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			token, err := extractBearer(c)
			if err != nil {
				return err
			}

			ctx := c.Request().Context()
			intoken, _, err := cli.OAuth2Api.IntrospectOAuth2Token(ctx).Token(token).Execute()
			if err != nil {
				log.Err(err).Msg("Failed to introspect token")
				return err
			}

			if !intoken.GetActive() {
				c.Response().Header().Set(echo.HeaderWWWAuthenticate, bearer+` error="invalid_token", error_description="The access token expired"`)
				return echo.ErrUnauthorized
			}

			c.Set(config.ContextKey, intoken)
			return next(c)
		}
	}
}

func extractBearer(c echo.Context) (string, error) {
	auth := c.Request().Header.Get(echo.HeaderAuthorization)

	l := len(bearer)
	if len(auth) > l+1 && strings.EqualFold(auth[:l], bearer) {
		token := auth[l+1:]
		return token, nil
	} else {
		c.Response().Header().Set(echo.HeaderWWWAuthenticate, bearer)
		return "", echo.ErrUnauthorized

	}
}
