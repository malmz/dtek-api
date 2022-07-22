package main

import (
	"os"

	"github.com/dtekcth/dtek-api/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04"})

	e := echo.New()

	e.HideBanner = true
	e.Debug = true

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:     true,
		LogMethod:  true,
		LogStatus:  true,
		LogLatency: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info().
				Str("method", v.Method).
				Str("uri", v.URI).
				Int("status", v.Status).
				Dur("latency", v.Latency).
				Msg("request")

			return nil
		},
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	{
		g := e.Group("/api")
		g.GET("/lunch", api.TodaysLunch)
	}
}
