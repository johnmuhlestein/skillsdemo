package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"skillsdemo/ent"
	"skillsdemo/ent/migrate"
	"skillsdemo/internal/surveyserver"
	"skillsdemo/rpc/survey"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	client := Open("postgresql://john:DonaldDuck@localhost:5433/postgres")

	defer client.Close()
    // Your code. For example:
    ctx := context.Background()
    if err := client.Debug().Schema.Create(ctx,migrate.WithDropIndex(true),migrate.WithDropColumn(true)); err != nil {
        log.Fatal(err)
    }
	log.Println("Successfully deployed database schema")

	server := &surveyserver.Server{}
	twirpHandler := survey.NewSurveyServer(server)

	http.ListenAndServe(":8080", twirpHandler)
}

// Open new connection
func Open(databaseUrl string) *ent.Client {
    db, err := sql.Open("pgx", databaseUrl)
    if err != nil {
        log.Fatal(err)
    }

    // Create an ent.Driver from `db`.
    drv := entsql.OpenDB(dialect.Postgres, db)
    return ent.NewClient(ent.Driver(drv))
}