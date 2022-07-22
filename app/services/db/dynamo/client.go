package dynamo

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	appConfig "github.com/davidenq/tweets-timeline-challenge/app/config"
)

func createLocalClient() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolver(
			aws.EndpointResolverFunc(
				func(service, region string) (aws.Endpoint, error) {
					return aws.Endpoint{URL: "http://localhost:8000"}, nil
				},
			),
		),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     "dummy",
				SecretAccessKey: "dummy",
				SessionToken:    "dummy",
				Source:          "Hard-coded credentials; values are irrelevant for local DynamoDB",
			},
		}),
	)
	if err != nil {
		log.Panicf("unable to load config, %v", err)
	}

	return cfg
}

func NewConn(appConfig appConfig.Config) dynamodb.Client {
	var cfg aws.Config
	var err error

	if appConfig.ENV == "local" {
		cfg = createLocalClient()
	} else {
		staticProvider := credentials.NewStaticCredentialsProvider(
			appConfig.AWSAccessKeyID,
			appConfig.AWSSecretAccessKey,
			"",
		)
		cfg, err = config.LoadDefaultConfig(
			context.Background(),
			config.WithCredentialsProvider(staticProvider),
			config.WithRegion(appConfig.AWSDefaultRegion),
		)
		if err != nil {
			log.Panicf("unable to load SDK config, %v", err)
		}
	}
	client := dynamodb.NewFromConfig(cfg)
	return *client
}
