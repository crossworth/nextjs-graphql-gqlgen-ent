package main

import (
	"context"
	"log"
	"net/http"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/crossworth/nextjs-graphql-gqlgen-ent/ent"
	"github.com/crossworth/nextjs-graphql-gqlgen-ent/ent/migrate"
	"github.com/crossworth/nextjs-graphql-gqlgen-ent/graph"
	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

func main() {
	client, err := ent.Open(dialect.SQLite, "file:posts?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
	}).Handler)

	srv := handler.NewDefaultServer(graph.NewSchema(client))
	srv.Use(entgql.Transactioner{
		TxOpener: client,
	})

	router.Handle("/", playground.Handler("Posts", "/query"))
	router.Handle("/query", srv)

	log.Println("listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("http server terminated", err)
	}
}
