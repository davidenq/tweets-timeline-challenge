package domain

type entityOptions struct {
	data      *[]byte
	checkData bool
}

type Option func(*entityOptions)

func WithData(data []byte) Option {
	return func(e *entityOptions) {
		e.data = &data
	}
}

func WithValidation(validation bool) Option {
	return func(e *entityOptions) {
		e.checkData = validation
	}
}

func NewEntity(entityName EntityName, options ...Option) (entity EntityPort, err error) {

	e := &entityOptions{}
	for _, option := range options {
		option(e)
	}

	s := schema{}
	s.get(entityName)

	if e.data != nil {
		s.fill(*e.data)
	}

	if e.checkData {
		s.validate()
	}

	if s.err != nil {
		return nil, s.err
	}

	return s.entity, nil
}
