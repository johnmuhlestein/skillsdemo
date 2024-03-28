package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"skillsdemo/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var EntClient *ent.Client

func init() {
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbDomain := os.Getenv("POSTGRES_DOMAIN")
	dbPort := os.Getenv("POSTGRES_PORT")
	db, err := sql.Open("pgx", fmt.Sprintf("postgresql://%s:%s@%s:%s/postgres",dbUser, dbPassword,dbDomain,dbPort))
    if err != nil {
        log.Fatal(err)
    }

    // Create an ent.Driver from `db`.
    drv := entsql.OpenDB(dialect.Postgres, db)
	Client := ent.NewClient(ent.Driver(drv))
	defer Client.Close()
    EntClient =  Client
}