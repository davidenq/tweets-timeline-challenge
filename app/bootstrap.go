package app

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/davidenq/tweets-timeline-challenge/app/config"
	"github.com/davidenq/tweets-timeline-challenge/app/domain"
	"github.com/davidenq/tweets-timeline-challenge/app/domain/usecases"
	"github.com/davidenq/tweets-timeline-challenge/app/services/db/dynamo"
	"github.com/davidenq/tweets-timeline-challenge/app/services/transport/http"
	"github.com/davidenq/tweets-timeline-challenge/app/ui/api"
	"github.com/davidenq/tweets-timeline-challenge/app/ui/api/handlers"
	"github.com/rs/zerolog/log"
)

type App struct {
	config          config.Config
	dbClient        dynamodb.Client
	httpClient      domain.HttpClientDrivenPort
	repository      domain.RepositoryDrivenPort
	timelineUsecase domain.TimeLineUsecaseFacadeDrivingPort
	userUsecase     domain.UserUsecaseDrivingPort
	handlers        handlers.HandlerDefinition
}

func (a *App) LoadConfig() *App {
	a.config = config.LoadConfig()
	return a
}

func (a *App) LoadServices() *App {
	a.httpClient = http.NewClient()
	a.dbClient = dynamo.NewConn(a.config)
	a.repository = dynamo.NewRepository(a.dbClient)
	return a
}

func (a *App) LoadDomain() *App {
	a.timelineUsecase = usecases.NewTimeLineUsecase(a.httpClient, a.repository, a.config)
	a.userUsecase = usecases.NewUserUsecase(a.repository, a.config)
	return a
}
func (a *App) LoadAPI() *App {
	a.handlers = handlers.NewHandlers(
		a.timelineUsecase,
		a.userUsecase,
	)
	return a
}
func (a *App) Init() {
	log.Logger = log.With().Caller().Logger()
	api.NewServer(a.config.APIPort, a.handlers)
}
