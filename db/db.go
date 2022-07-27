package db

import (
	"context"
	"fmt"
	"os"

	"github.com/dtekcth/dtek-api/ent"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

var db *ent.Client

func Init() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, user, dbname, password))
	if err != nil {
		log.Fatal().Err(err).Msg("failed opening connection to database")
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("failed creating schema resources")
	}

	db = client
}

func Get() *ent.Client {
	return db
}

func Set(client *ent.Client) {
	db = client
}

func Close() {
	db.Close()
	db = nil
}
