package domain

import (
	"testing"

	"github.com/davidenq/tweets-timeline-challenge/app/domain/schemas"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type testEntityFactory struct {
	description string
	entityName  EntityName
	options     []Option

	expectedEntity EntityPort
	expectedError  bool
}

var (
	uuidv4           = uuid.New()
	oauthFakeData    = `{"id": "` + uuidv4.String() + `"}`
	badOauthFakeData = `{"id": "nonuuid"}`
)

func TestWithData(t *testing.T) {
	t.Run("should return entityOptions struct set with data", func(t *testing.T) {
		e := &entityOptions{}
		data := []byte(oauthFakeData)
		f := WithData(data)
		f(e)
		assert.Equal(t, &data, e.data)
	})
}

func TestWithValidation(t *testing.T) {
	t.Run("should return entityOptions struct set with checkData ", func(t *testing.T) {
		e := &entityOptions{}
		expected := true
		f := WithValidation(expected)
		f(e)
		assert.Equal(t, expected, e.checkData)
	})
}

func TestNewEntity(t *testing.T) {

	tests := []testEntityFactory{
		{
			description:    "should return nil schema and error generated by empty schema name",
			entityName:     "",
			expectedEntity: nil,
			expectedError:  true,
		},
		{
			description:    "should return an empty oauth schema and nil error",
			entityName:     OAuth,
			expectedEntity: &schemas.OAuthSchema{},
			expectedError:  false,
		},
		{
			description: "should return a schema with data and nil error",
			entityName:  OAuth,
			options: []Option{
				WithData([]byte(oauthFakeData)),
			},
			expectedEntity: &schemas.OAuthSchema{
				ID: uuidv4.String(),
			},
			expectedError: false,
		},
		{
			description: "should return a nil schema and error generated by validation schema",
			entityName:  OAuth,
			options: []Option{
				WithData([]byte(badOauthFakeData)),
				WithValidation(true),
			},
			expectedEntity: nil,
			expectedError:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			var actualEntity EntityPort
			var actualError error
			if test.options != nil {
				actualEntity, actualError = NewEntity(test.entityName, test.options...)
			} else {
				actualEntity, actualError = NewEntity(test.entityName)
			}
			assert.Equal(t, test.expectedEntity, actualEntity)
			assert.Equal(t, test.expectedError, actualError != nil)
		})
	}
}
