package domain

import (
	"encoding/json"
	"errors"

	"github.com/davidenq/tweets-timeline-challenge/app/domain/schemas"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

type Schema struct {
	Entity EntityPort
	err    error
}

func (s *Schema) Get(entityName EntityName) *Schema {
	switch EntityName(entityName) {
	case OAuth:
		s.Entity = &schemas.OAuthSchema{}
	case Timelines:
		s.Entity = &schemas.TimelinesSchema{}
	case User:
		s.Entity = &schemas.UserSchema{}
	default:
		s.err = errors.New(SchemaNotFound)
	}
	return s
}

func (s *Schema) Fill(data []byte) *Schema {
	if s.Entity == nil {
		s.err = errors.New(NilSchema)
		log.Error().Err(s.err)
		return s
	}

	if len(data) == 0 {
		s.err = errors.New(DataNotBeEmpty)
		log.Error().Err(s.err)
		return s
	}
	err := json.Unmarshal(data, s.Entity)
	if err != nil {
		log.Error().Err(s.err)
		s.err = errors.New(err.Error())
	}
	return s
}

func (s *Schema) Validate() *Schema {
	if s.Entity == nil {
		log.Error().Err(s.err)
		s.err = errors.New(NilSchema)
		return s
	}
	validate := validator.New()
	err := s.Entity.Validate(*validate)
	if err != nil {
		log.Error().Err(s.err)
		s.err = err
		return s
	}

	return s
}
