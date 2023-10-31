package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Purvig648/graphql-demo/database"
	"github.com/Purvig648/graphql-demo/graph"
	"github.com/Purvig648/graphql-demo/repository"
	"github.com/Purvig648/graphql-demo/service"
	"github.com/rs/zerolog/log"
)

const defaultPort = "8080"

func main() {
	svc, err := StartApp()
	if err != nil {
		log.Info().Err(err).Msg("could not startapp")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Service: svc,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal().Err(http.ListenAndServe(":"+port, nil))

}

func StartApp() (service.UserService, error) {
	db, err := database.Open()
	if err != nil {
		return &service.Service{}, fmt.Errorf("connecting to database %w", err)
	}
	pg, err := db.DB()
	if err != nil {
		return &service.Service{}, fmt.Errorf("failed to get database instance: %w ", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = pg.PingContext(ctx)
	if err != nil {
		return &service.Service{}, fmt.Errorf("database is not connected: %w ", err)
	}
	repo, err := repository.NewRepository(db)
	if err != nil {
		return &service.Service{}, fmt.Errorf("could not initialize repo layer: %w ", err)
	}
	svc, err := service.NewService(repo)
	if err != nil {
		return &service.Service{}, fmt.Errorf("could not initialize service layer: %w ", err)
	}
	return svc, nil
}
