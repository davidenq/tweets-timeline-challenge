package main

import (
	"context"

	"github.com/davidenq/tweets-timeline-challenge/app/config"
	"github.com/davidenq/tweets-timeline-challenge/app/domain/usecases"
	"github.com/davidenq/tweets-timeline-challenge/app/services/db/dynamo"
	"github.com/davidenq/tweets-timeline-challenge/app/services/transport/http"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg := config.LoadConfig()
	runMigration(cfg)
	generateToken(cfg)
}

func generateToken(cfg config.Config) {

	dbClient := dynamo.NewConn(cfg)
	httpClient := http.NewClient()
	repository := dynamo.NewRepository(dbClient)
	oauthUsecase := usecases.NewOAuth(cfg, httpClient, repository)
	_, err := oauthUsecase.GetToken(context.TODO())
	if err != nil {
		log.Error().Msg(err.Error())
		panic(err)
	}
	log.Info().Msg("access token to connect with twitter has been requested successful!")
}

func runMigration(cfg config.Config) {
	dbClient := dynamo.NewConn(cfg)
	repository := dynamo.NewRepository(dbClient)
	migrate := usecases.NewMigrate(repository)
	log.Info().Msg("tables will be create on " + cfg.AWSDefaultRegion + " region")
	migrate.DeleteTables()
	migrate.CreateTables()
}
