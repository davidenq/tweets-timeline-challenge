package usecases

import (
	"context"

	"github.com/davidenq/tweets-timeline-challenge/app/config"
	"github.com/davidenq/tweets-timeline-challenge/app/domain"
	"github.com/davidenq/tweets-timeline-challenge/app/utils/errorkit"
	"github.com/rs/zerolog/log"
)

type UserUsecase struct {
	dbClient domain.RepositoryDrivenPort
	config   config.Config
}

func (UserUsecase) CheckSession(ctx context.Context, username string) error {
	errMessage := "not implemented yet"
	log.Logger = log.With().Caller().Logger()
	err := errorkit.NewServerError(errMessage)
	log.Error().Err(err).Msg(string(errMessage))
	return err
}
func (UserUsecase) Update(ctx context.Context, username string) error {
	errMessage := "not implemented yet"

	err := errorkit.NewServerError(errMessage)
	log.Error().Err(err).Msg(string(errMessage))
	return err
}

func NewUserUsecase(
	dbClient domain.RepositoryDrivenPort,
	config config.Config,
) domain.UserUsecaseDrivingPort {
	return &UserUsecase{
		dbClient: dbClient,
		config:   config,
	}
}
