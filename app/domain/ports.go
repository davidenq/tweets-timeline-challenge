package domain

import (
	"context"
	"net/http"

	"github.com/davidenq/tweets-timeline-challenge/app/domain/schemas"
	transport "github.com/davidenq/tweets-timeline-challenge/app/services/transport/http"

	"github.com/go-playground/validator/v10"
)

//driving ports
type OAuthUsecaseFacadeDrivingPort interface {
	GetToken(ctx context.Context) (token *string, err error)
	SaveInDB(oauthSchema schemas.OAuthSchema) error
}

type TimeLineUsecaseFacadeDrivingPort interface {
	GetTimeline(ctx context.Context, username string, timelimeElements string) (*schemas.TimelinesSchema, error)
}

type UserUsecaseDrivingPort interface {
	CheckSession(ctx context.Context, username string) error
	Update(ctx context.Context, username string) error
}

type MigrationUsecaseFacadeDrivingPort interface {
	CreateTables() error
	DeleteTables() error
}

//domain ports
type EntityPort interface {
	Validate(validate validator.Validate) error
}

//driven ports
type RepositoryDrivenPort interface {
	GetLastRecord(tableName string, out interface{}) error
	GetEntityByID(tableName string, id string, out interface{}) error
	GetEntityByUsername(tableName string, username string, out interface{}) error
	SaveEntity(tableName string, entity interface{}) error
	EditEntity(tableName string, entity map[string]interface{}) error
	DeleteEntity(tableName string, id string) error
	CreateTables(interface{}) error
	RemoveTables() error
}

type HttpClientDrivenPort interface {
	Config(r transport.Request) *transport.Request
	Do(ctx context.Context, dataOutcome interface{}) (*http.Response, error)
}
