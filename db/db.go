package db

import (
	"context"
	"os"

	"entgo.io/ent/dialect"
	"github.com/dtekcth/dtek-api/ent"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
)

var db *ent.Client

func Init() {

	dbType := os.Getenv("DB_TYPE")
	if dbType == "" {
		dbType = dialect.SQLite
	}

	dbDsn := os.Getenv("DB_DSN")
	if dbDsn == "" {
		dbDsn = "file:debug.db?_fk=1"
	}

	client, err := ent.Open(dbType, dbDsn)
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
