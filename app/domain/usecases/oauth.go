package usecases

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/davidenq/tweets-timeline-challenge/app/config"
	"github.com/davidenq/tweets-timeline-challenge/app/domain"
	"github.com/davidenq/tweets-timeline-challenge/app/domain/schemas"
	transport "github.com/davidenq/tweets-timeline-challenge/app/services/transport/http"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type OAuthUsecaseFacade struct {
	httpClient domain.HttpClientDrivenPort
	dbClient   domain.RepositoryDrivenPort
	config     config.Config
}

func (usecase OAuthUsecaseFacade) GetToken(ctx context.Context) (token *string, err error) {
	ptrToken, err := usecase.getFromStorage(ctx)
	if err != nil {
		log.Error().Stack().Err(err).Msg("")
		return nil, err
	}
	if *ptrToken != "" {
		return ptrToken, nil
	} else {
		return usecase.fetchFromService(ctx)
	}
}

func (usecase OAuthUsecaseFacade) getFromStorage(ctx context.Context) (token *string, err error) {
	return getLastToken(usecase.dbClient, ctx)
}

func (usecase OAuthUsecaseFacade) fetchFromService(ctx context.Context) (token *string, err error) {
	var oauthSchema schemas.OAuthSchema
	oauthSchema.ID = uuid.New().String()
	req := transport.Request{
		URL:    usecase.config.TwitterOAuthURL,
		Method: http.MethodPost,
		Headers: map[string]string{
			"Authorization": "Basic " + basicAuth(usecase.config.TwitterAPIKey, usecase.config.TwitterAPISecretKey),
		},
	}
	response, err := usecase.httpClient.Config(req).Do(ctx, &oauthSchema)
	if err != nil {
		log.Error().Stack().Err(err).Msg("")
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, err
	}
	oauthSchema.ExpiresOn = "300" //5 minutes to expires
	usecase.SaveInDB(oauthSchema)
	return &oauthSchema.AccessToken, nil
}

func (usecase OAuthUsecaseFacade) SaveInDB(oauthSchema schemas.OAuthSchema) error {
	var data map[string]interface{}
	bytes, err := json.Marshal(oauthSchema)
	if err != nil {
		log.Error().Stack().Err(err).Msg("")
		return err
	}
	json.Unmarshal(bytes, &data)
	err = usecase.dbClient.SaveEntity(string(domain.OAuth), data)
	if err != nil {
		log.Error().Stack().Err(err).Msg("")
		return err
	}
	return nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func getLastToken(client domain.RepositoryDrivenPort, ctx context.Context) (token *string, err error) {
	var oauthSchema schemas.OAuthSchema
	err = client.GetLastRecord(string(domain.OAuth), &oauthSchema)
	if err != nil {
		log.Error().Stack().Err(err).Msg("")
		return nil, err
	}
	return &oauthSchema.AccessToken, nil
}

func NewOAuth(
	config config.Config,
	httpClient domain.HttpClientDrivenPort,
	dbClient domain.RepositoryDrivenPort,
) domain.OAuthUsecaseFacadeDrivingPort {
	return &OAuthUsecaseFacade{
		httpClient: httpClient,
		dbClient:   dbClient,
		config:     config,
	}
}
