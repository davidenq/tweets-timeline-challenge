package usecases

import (
	"reflect"
	"strings"
	"time"

	"github.com/davidenq/tweets-timeline-challenge/app/domain"
	"github.com/rs/zerolog/log"
)

type Attribute struct {
	Name string
	Type string
}

type SchemaToTable struct {
	Key        string
	Attributes []Attribute
	TableName  string
}

type MigrateUsecase struct {
	repository domain.RepositoryDrivenPort
}

func (m MigrateUsecase) CreateTables() error {
	oauthMigration := &SchemaToTable{}
	err := m.repository.CreateTables(*oauthMigration.generateTableDefinition(domain.OAuth))
	if err != nil {
		log.Error().Stack().Err(err).Msg("")
		return err
	}
	userMigration := &SchemaToTable{}
	err = m.repository.CreateTables(*userMigration.generateTableDefinition((domain.User)))
	if err != nil {
		log.Error().Stack().Err(err).Msg("")
		return err
	}
	time.Sleep(2 * time.Second)
	return nil
}

func (m MigrateUsecase) DeleteTables() error {
	return m.repository.RemoveTables()
}

func (st *SchemaToTable) generateTableDefinition(entityName domain.EntityName) *SchemaToTable {
	attributes := make([]Attribute, 0)

	schema := domain.Schema{}
	s := schema.Get(entityName)
	e := reflect.ValueOf(s.Entity).Elem()

	for i := 0; i < e.NumField(); i++ {
		tag := strings.Split(e.Type().Field(i).Tag.Get("dynamo"), ":")
		if len(tag) > 1 {
			if tag[0] == "key" {
				st.Key = tag[1]
				attribute := Attribute{
					Name: tag[1],
					Type: e.Type().Field(i).Type.String(),
				}
				attributes = append(attributes, attribute)
			}

		}
	}
	st.Attributes = attributes
	st.TableName = string(entityName)
	return st
}

func NewMigrate(repository domain.RepositoryDrivenPort) domain.MigrationUsecaseFacadeDrivingPort {
	return MigrateUsecase{
		repository: repository,
	}
}
