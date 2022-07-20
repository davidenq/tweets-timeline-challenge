package usecases

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/davidenq/tweets-timeline-challenge/app/config"
	"github.com/davidenq/tweets-timeline-challenge/app/domain"
	"github.com/davidenq/tweets-timeline-challenge/app/domain/schemas"
	"github.com/davidenq/tweets-timeline-challenge/mocks"
	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
)

func TestSaveOAuthSchemaInDBMock(t *testing.T) {

	type argsForDBClientMock struct {
		tableName   domain.EntityName
		oauthSchema schemas.OAuthSchema

		decidedError error
	}

	type testDef struct {
		description string
		schema      schemas.OAuthSchema

		expectedError bool

		*argsForDBClientMock
	}
	oauthSchema := schemas.OAuthSchema{
		ID:          uuid.New().String(),
		AccessToken: uuid.New().String(),
		TokenType:   "bearer_token",
	}
	tests := []testDef{
		{
			description:   "should return error generated by dbClient.SaveEntity usecase",
			schema:        oauthSchema,
			expectedError: true,
			argsForDBClientMock: &argsForDBClientMock{
				tableName:    domain.OAuth,
				oauthSchema:  oauthSchema,
				decidedError: errors.New("error"),
			},
		},
		{
			description:   "should return nil error when all is ok",
			schema:        oauthSchema,
			expectedError: false,
			argsForDBClientMock: &argsForDBClientMock{
				tableName:    domain.OAuth,
				oauthSchema:  oauthSchema,
				decidedError: nil,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {

			httpClient := &mocks.HttpClientDrivenPort{}
			dbClient := &mocks.RepositoryDrivenPort{}
			config := config.Config{}
			oauthUsecase := NewOAuth(config, httpClient, dbClient)
			var oauthMap map[string]interface{}
			bytes, _ := json.Marshal(test.schema)
			json.Unmarshal(bytes, &oauthMap)
			dbClient.
				On(
					"SaveEntity",
					string(domain.OAuth),
					oauthMap,
				).
				Return(test.argsForDBClientMock.decidedError)

			actualError := oauthUsecase.SaveInDB(test.schema)
			assert.Equal(t, test.argsForDBClientMock.decidedError != nil, actualError != nil)
		})

	}

}
