package integration

import (
	"context"
	"testing"
	"time"

	"github.com/davidenq/tweets-timeline-challenge/app/config"
	"github.com/davidenq/tweets-timeline-challenge/app/domain/usecases"
	"github.com/davidenq/tweets-timeline-challenge/app/services/db/dynamo"
	"github.com/davidenq/tweets-timeline-challenge/app/services/transport/http"
	"github.com/stretchr/testify/assert"
)

func TestDynamoDBTwitterIntegration(t *testing.T) {
	cfg := config.LoadConfig()
	dbConn := dynamo.NewConn(cfg)
	httpClient := http.NewClient()
	repository := dynamo.NewRepository(dbConn)
	migrateUsecases := usecases.NewMigrate(repository)
	oauthUsecase := usecases.NewOAuth(cfg, httpClient, repository)

	t.Run("should create tables in DynamoDB service", func(t *testing.T) {
		migrateUsecases.DeleteTables()
		actualErr := migrateUsecases.CreateTables()
		assert.NoError(t, actualErr)
	})

	t.Run("should get a token from Twitter service and store in DynamoDB", func(t *testing.T) {
		time.Sleep(10 * time.Second)
		token, actualErr := oauthUsecase.GetToken(context.TODO())
		assert.NotNil(t, token)
		assert.NoError(t, actualErr)
	})

	t.Run("should destroy tables in DynamoDB service", func(t *testing.T) {
		actualErr := migrateUsecases.DeleteTables()
		assert.NoError(t, actualErr)
	})
}
