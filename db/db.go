package db

import (
	"context"

	"github.com/dtekcth/dtek-api/ent"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
)

var db *ent.Client

func Init() {
	client, err := ent.Open("sqlite3", "file:test.db?_fk=1")
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
