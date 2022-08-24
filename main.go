package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dtekcth/dtek-api/api"
	"github.com/dtekcth/dtek-api/db"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Validator struct {
	validate *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	if err := v.validate.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04"})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	db.Init()
	defer db.Close()

	e := echo.New()

	e.HideBanner = true
	e.Debug = true
	e.Validator = &Validator{validate: validator.New()}

	log.Info().Msg("Starting server")
	log.Debug().Msg("Debug mode enabled")

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
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			log.Error().
				Err(err).
				Str("stack", string(stack)).
				Msg("panic")
			return nil
		},
	}))
	e.Use(middleware.CORS())

	{
		g := e.Group("/api")
		g.GET("/lunch", api.GetLunch)

		g.GET("/news", api.GetAllNews)
		g.POST("/news", api.CreateNews)
		g.GET("/news/:id", api.GetNews)
		g.PUT("/news/:id", api.UpdateNews)
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal().Err(e.Start(fmt.Sprintf("%s:%s", host, port))).Msg("server exited")
}
