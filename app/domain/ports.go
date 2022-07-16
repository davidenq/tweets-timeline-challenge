package domain

import (
	"github.com/go-playground/validator/v10"
)

type EntityPort interface {
	Validate(validate validator.Validate) error
}
