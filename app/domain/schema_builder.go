package domain

import (
	"encoding/json"
	"errors"

	"github.com/davidenq/tweets-timeline-challenge/app/domain/schemas"
	"github.com/go-playground/validator/v10"
)

type schema struct {
	entity EntityPort
	err    error
}

func (s *schema) get(entityName EntityName) *schema {
	switch EntityName(entityName) {
	case OAuth:
		s.entity = &schemas.OAuthSchema{}
	case Profile:
		s.entity = &schemas.ProfileSchema{}
	case Tweets:
		s.entity = &schemas.TimelinesSchema{}
	case User:
		s.entity = &schemas.UserSchema{}
	default:
		s.err = errors.New(SchemaNotFound)
	}
	return s
}

func (s *schema) fill(data []byte) *schema {
	if s.entity == nil {
		s.err = errors.New(NilSchema)
		return s
	}

	if len(data) == 0 {
		s.err = errors.New(DataNotBeEmpty)
		return s
	}
	err := json.Unmarshal(data, s.entity)
	if err != nil {
		s.err = errors.New(err.Error())
	}
	return s
}

func (s *schema) validate() *schema {
	if s.entity == nil {
		s.err = errors.New(NilSchema)
		return s
	}
	validate := validator.New()
	err := s.entity.Validate(*validate)
	if err != nil {
		s.err = err
		return s
	}

	return s
}
