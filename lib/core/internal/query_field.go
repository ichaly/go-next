package internal

import "github.com/graphql-go/graphql"

type RootField struct {
	Object *graphql.Object
}

func (my *RootField) Name() string {
	return my.Object.Name()
}

func (my *RootField) Field() *graphql.Field {
	return &graphql.Field{
		Name: my.Name(),
		Type: my.Object,
	}
}
