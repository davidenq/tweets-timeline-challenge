package schemas

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type ProfileSchema struct {
	ID        string    `json:"id" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (s *ProfileSchema) Validate(validate validator.Validate) error {
	return nil
}
