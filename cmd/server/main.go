package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"skillsdemo/ent/migrate"
	"skillsdemo/internal/database"
	"skillsdemo/internal/hooks"
	"skillsdemo/internal/surveyserver"
	"skillsdemo/rpc/survey"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {

	// Your code. For example:
    ctx := context.Background()
    if err := database.EntClient.Debug().Schema.Create(ctx,migrate.WithDropIndex(true),migrate.WithDropColumn(true)); err != nil {
        log.Fatal(err)
    }
	log.Println("Successfully deployed database schema")
	hook := hooks.LogginHooks(os.Stderr)
	server := &surveyserver.Server{}
	twirpHandler := survey.NewSurveyServer(server,hook)

	http.ListenAndServe(":8080", twirpHandler)
}

// Open new connection
/*
func Open(databaseUrl string) *ent.Client {
    db, err := sql.Open("pgx", databaseUrl)
    if err != nil {
        log.Fatal(err)
    }

    // Create an ent.Driver from `db`.
    drv := entsql.OpenDB(dialect.Postgres, db)
    return ent.NewClient(ent.Driver(drv))
}
*/