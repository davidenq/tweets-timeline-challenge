package usecases

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/davidenq/tweets-timeline-challenge/app/config"
	"github.com/davidenq/tweets-timeline-challenge/app/domain"
	"github.com/davidenq/tweets-timeline-challenge/app/domain/schemas"
	transport "github.com/davidenq/tweets-timeline-challenge/app/services/transport/http"
	"github.com/davidenq/tweets-timeline-challenge/app/utils/errorkit"
	"github.com/rs/zerolog/log"
)

var (
	usernamePath  string = "/2/users/by/username/%s"
	timelinesPath string = "/1.1/statuses/user_timeline.json"

	queryParamUsername map[string]string = map[string]string{
		"user.fields": "profile_image_url,description,name",
	}
)

type TimeLineUsecaseFacade struct {
	httpClient domain.HttpClientDrivenPort
	dbClient   domain.RepositoryDrivenPort
	config     config.Config
}

func (usecase TimeLineUsecaseFacade) GetTimeline(ctx context.Context, username string, timelimeElements string) (*schemas.TimelinesSchema, error) {
	var user *schemas.UserSchema
	var timeline *schemas.TimelinesSchema
	var userInMap map[string]string
	var err error
	user, err = usecase.getFromStorage(ctx, username)
	if err != nil {
		log.Error().Stack().Err(err).Msg("")
		return nil, err
	}
	if user.Username == "" {
		//get user from service
		user, userInMap, err = usecase.fetchUserByUsernameFromService(ctx, username)
		if err != nil {
			log.Error().Stack().Err(err).Msg("")
			return nil, err
		}
		//record user into db
		err = usecase.saveInDB(userInMap)
		if err != nil {
			log.Error().Stack().Err(err).Msg("")
			return nil, err
		}
	}
	timeline, err = usecase.fetchTimelinesByUserIDFromService(ctx, user.ID, timelimeElements)
	if err != nil {
		log.Error().Stack().Err(err).Msg("")
		return nil, err
	}
	return timeline, nil
}

func (usecase TimeLineUsecaseFacade) getFromStorage(ctx context.Context, username string) (*schemas.UserSchema, error) {
	var userSchema schemas.UserSchema
	err := usecase.dbClient.GetEntityByUsername(string(domain.User), username, &userSchema)
	if err != nil {
		log.Error().Stack().Err(err).Msg("")
		return nil, err
	}
	return &userSchema, nil
}

func (usecase TimeLineUsecaseFacade) fetchUserByUsernameFromService(ctx context.Context, username string) (*schemas.UserSchema, map[string]string, error) {
	var data map[string]map[string]string
	err := usecase.fetchData(ctx, http.MethodGet, usecase.config.TwitterUsersURL, fmt.Sprintf(usernamePath, username), queryParamUsername, &data)
	if err != nil {
		log.Error().Stack().Err(err).Msg("")
		return nil, nil, err
	}
	userSchema := schemas.UserSchema{
		ID:              data["data"]["id"],
		Username:        data["data"]["username"],
		Name:            data["data"]["name"],
		Description:     data["data"]["description"],
		ProfileImageURL: data["data"]["description"],
	}
	return &userSchema, data["data"], nil
}

func (usecase TimeLineUsecaseFacade) fetchTimelinesByUserIDFromService(ctx context.Context, id, count string) (*schemas.TimelinesSchema, error) {
	var timeline schemas.TimelinesSchema
	queryParamsTimelines := map[string]string{
		"user_id": id,
		"count":   count,
	}
	err := usecase.fetchData(ctx, http.MethodGet, usecase.config.TwitterUsersURL, timelinesPath, queryParamsTimelines, &timeline)
	if err != nil {
		log.Error().Stack().Err(err).Msg("")
		return nil, err
	}
	return &timeline, nil
}

func (usecase TimeLineUsecaseFacade) saveInDB(data map[string]string) error {
	err := usecase.dbClient.SaveEntity(string(domain.User), data)
	if err != nil {
		log.Error().Stack().Err(err).Msg("")
		return err
	}
	return nil
}

func (usecase TimeLineUsecaseFacade) fetchData(ctx context.Context, method, baseURL, path string, params map[string]string, out interface{}) error {
	token, err := getLastToken(usecase.dbClient, ctx)

	if err != nil {
		log.Error().Stack().Err(err).Msg("")
		return err
	}
	url, _ := generateURL(baseURL, path, params)
	req := transport.Request{
		URL:    url.String(),
		Method: method,
		Headers: map[string]string{
			"Authorization": "Bearer " + *token,
		},
	}
	response, err := usecase.httpClient.Config(req).Do(ctx, &out)
	if err != nil {
		log.Error().Stack().Err(err).Msg("")
		return err
	}
	if response.StatusCode != http.StatusOK {
		log.Error().Stack().Err(err).Msg("")
		return err
	}
	return nil
}

func generateURL(rawURL, path string, params map[string]string) (*url.URL, error) {
	baseURL, err := url.Parse(rawURL)
	if err != nil {
		log.Error().Stack().Err(err).Msg("")
		return nil, errorkit.NewServerError(err.Error())
	}
	baseURL.Path = path
	values := url.Values{}
	for k, v := range params {
		values.Add(k, v)
	}
	baseURL.RawQuery = values.Encode()
	return baseURL, nil
}

func NewTimeLineUsecase(
	httpClient domain.HttpClientDrivenPort,
	dbClient domain.RepositoryDrivenPort,
	config config.Config,
) domain.TimeLineUsecaseFacadeDrivingPort {
	return &TimeLineUsecaseFacade{
		httpClient: httpClient,
		dbClient:   dbClient,
		config:     config,
	}
}
