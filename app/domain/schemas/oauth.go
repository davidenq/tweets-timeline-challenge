package schemas

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type OAuthSchema struct {
	ID        string    `json:"id" validate:"required,uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (s *OAuthSchema) Validate(validate validator.Validate) error {
	err := validate.Struct(s)
	if err != nil {
		return err
	}
	return nil
}
