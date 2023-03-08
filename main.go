package main

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"net/http"
	"os"

	"github.com/dtekcth/dtek-api/db"
	"github.com/dtekcth/dtek-api/handler"
	mw "github.com/dtekcth/dtek-api/middleware"
	"github.com/dtekcth/dtek-api/service/lunch"
	"github.com/dtekcth/dtek-api/service/pr"
	"github.com/pressly/goose/v3"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	// Init logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04"})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal().Err(err).Msg("Failed to load .env file")
	}

	runMigrations()
	pool, db := createDb()
	defer pool.Close()

	/* 	db, err := ent.Open(os.Getenv("DB_TYPE"), os.Getenv("DB_DSN"))
	   	if err != nil {
	   		log.Fatal().Err(err).Msg("failed opening connection to database")
	   	}

	   	// Run migrations
	   	if err := db.Schema.Create(context.Background()); err != nil {
	   		log.Fatal().Err(err).Msg("failed creating schema resources")
	   	}

	   	defer db.Close() */

	e := createEcho()

	authmw := mw.Hydra()
	lunchService := &lunch.Service{Db: db}
	prService := pr.NewService("c_012de84a3e299139f30e87ac76c2704f1585161d031007c5b81438af0e7f5b15@group.calendar.google.com")
	env := &handler.Env{Db: db, LunchService: lunchService, PrService: prService}

	{
		g := e.Group("/api")
		g.GET("/lunch", env.GetLunch)

		g.GET("/news", env.ListNews)
		g.POST("/news", env.CreateNews, authmw)
		g.GET("/news/:id", env.GetNews)
		g.PUT("/news/:id", env.UpdateNews, authmw)

		g.POST("/pr", env.CreateEvent)
	}

	e.GET("/hydra", env.HydraTest, authmw)

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.
		Fatal().
		Err(
			e.Start(fmt.Sprintf("%s:%s", host, port)),
		).
		Msg("server exited")
}

//go:embed db/schema/*.sql
var embededMigrations embed.FS

func runMigrations() {
	db, err := sql.Open("pgx", os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database for migrations")
	}
	defer db.Close()

	goose.SetBaseFS(embededMigrations)
	goose.SetDialect("pgx")
	if err = goose.Up(db, "db/schema"); err != nil {
		log.Fatal().Err(err).Msg("failed to run migrations")
	}
}

func createDb() (*pgxpool.Pool, *db.Queries) {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}
	db := db.New(pool)
	return pool, db
}

func createEcho() *echo.Echo {
	e := echo.New()

	//e.HideBanner = true
	e.Debug = true
	e.Validator = &Validator{validate: validator.New()}

	// Set echo to use zerolog
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

	// Recover from panics in handlers
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			log.Error().
				Err(err).
				Str("stack", string(stack)).
				Msg("panic")
			return nil
		},
	}))

	// Add CORS
	e.Use(middleware.CORS())

	return e
}

type Validator struct {
	validate *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	if err := v.validate.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
