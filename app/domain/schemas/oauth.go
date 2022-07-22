package schemas

import (
	"github.com/go-playground/validator/v10"
)

type OAuthSchema struct {
	ID          string `json:"id" validate:"required,uuid" dynamo:"key:id"`
	TokenType   string `json:"token_type" dynamo:"attribute:token_type"`
	AccessToken string `json:"access_token" dynamo:"attribute:access_token"`
	ExpiresOn   string `json:"expires_on" dynamo:"attribute:expires_on"`
}

func (s *OAuthSchema) Validate(validate validator.Validate) error {
	err := validate.Struct(s)
	if err != nil {
		return err
	}
	return nil
}
