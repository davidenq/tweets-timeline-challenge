package schemas

import (
	"github.com/go-playground/validator/v10"
)

type UserSchema struct {
	ID              string `json:"id" validate:"required" dynamo:"attribute:id"`
	Name            string `json:"name" validate:"required" dynamo:"attribute:name"`
	Username        string `json:"username" validate:"required" dynamo:"key:username"`
	FullName        string `json:"full_name" validate:"required" dynamo:"attribute:fullname"`
	Description     string `json:"description" validate:"required" dynamo:"attribute:description"`
	ProfileImageURL string `json:"profile_image_url" dynamo:"attribute:profile_image_url"`
}

func (s *UserSchema) Validate(validate validator.Validate) error {
	return nil
}
