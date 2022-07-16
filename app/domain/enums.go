package domain

type EntityName string

const (
	OAuth   EntityName = "oauth"
	Profile EntityName = "profile"
	Tweets  EntityName = "tweets"
	User    EntityName = "user"
)

const (
	SchemaNotFound = "schema not found"
	NilSchema      = "a schema must be generated previously"
	DataNotBeEmpty = "data must not be empty"
)
