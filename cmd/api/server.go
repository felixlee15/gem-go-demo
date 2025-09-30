package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"

	"go-demo/config"
	"go-demo/ent"
	"go-demo/internal/app/delivery/graph"
	"go-demo/internal/pkg/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	log.Println("Starting the application...")
	executeCommand()
}

// This function decides which server, worker, or scheduler to start based on the task parameter provided when running the program.
// Example, no use
func executeCommand() {
	runServer()
}

func runServer() {
	dbConfig := config.GetDBConfig()
	db, err := sql.Open("postgres", dbConfig.DSN())
	if err != nil {
		log.Fatal(err)
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal(err)
	}

	resolver := &graph.Resolver{Client: client}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", srv)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
