package config

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	ENV     string `env:"ENV"`
	APIPort string `env:"API_PORT"`

	DynamoDBDomain string `env:"DYNAMODB_DOMAIN"`
	DynamoDBPort   string `env:"DYNAMODB_PORT"`

	AWSAccessKeyID     string `env:"AWS_ACCESS_KEY_ID"`
	AWSSecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY"`
	AWSDefaultRegion   string `env:"AWS_DEFAULT_REGION"`

	TwitterAPIKey       string `env:"TWITTER_API_KEY"`
	TwitterAPISecretKey string `env:"TWITTER_API_SECRET_KEY"`
	TwitterOAuthURL     string `env:"TWITTER_OAUTH_URL"`
	TwitterUsersURL     string `env:"TWITTER_USERS_URL"`
}

func LoadConfig() Config {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Panicf("%+v\n", err)
	}
	return cfg
}
