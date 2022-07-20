package domain

type EntityName string

const (
	OAuth     EntityName = "oauth"
	Profile   EntityName = "profile"
	Timelines EntityName = "timelines"
	User      EntityName = "user"
)

const (
	SchemaNotFound = "schema not found"
	NilSchema      = "a schema must be generated previously"
	DataNotBeEmpty = "data must not be empty"
)
