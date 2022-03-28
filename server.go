package main

import (
	"log"
	"movie_graphql_be/graph"
	"movie_graphql_be/graph/generated"
	"movie_graphql_be/internal/auth"
	"movie_graphql_be/pkg/db/mysql"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)
const defaultHost = "localhost"
const defaultPort = "8080"

func main() {
	godotenv.Load()
	host := os.Getenv("CONN_HOST")
	port := os.Getenv("CONN_PORT")
	if host == "" {
		host = defaultHost
	}
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(auth.JwtMiddleware())
	mysql.InitDB()
	mysql.Migrate()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://%s:%s/ for GraphQL playground", host, port)
	log.Fatal(http.ListenAndServe(host+":"+port, router))
}
