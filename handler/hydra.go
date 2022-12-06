package handler

import (
	"fmt"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	hydra "github.com/ory/hydra-client-go"
	"github.com/rs/zerolog/log"
)

func extractBearer(c echo.Context) (string, error) {
	parts := strings.Split(c.Request().Header.Get("Authorization"), " ")
	if len(parts) == 2 && strings.EqualFold(parts[0], "Bearer") {
		return parts[1], nil
	} else {
		return "", fmt.Errorf("No bearer token found")
	}
}

var cli *hydra.APIClient

func InitHydra() {
	conf := hydra.NewConfiguration()
	conf.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %v", os.Getenv("ORY_PAT")))
	conf.Servers = []hydra.ServerConfiguration{
		{
			URL: os.Getenv("ORY_URL"),
		},
	}
	cli = hydra.NewAPIClient(conf)
}

// echo request handler
func (e *Env) HydraTest(c echo.Context) error {
	token := c.Get("user").(*hydra.IntrospectedOAuth2Token)
	log.Printf("token %#v", token)
	println("sub", token.GetSub())
	println("username", token.GetUsername())
	return c.String(200, "Hello, World!")
}
