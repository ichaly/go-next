package gql

import (
	"github.com/graphql-go/graphql"
	"reflect"
)

type GqlID interface {
	ID()
}

func (my *Engine) asId(typ reflect.Type) (graphql.Type, error) {
	if _, isId := reflect.New(typ).Interface().(GqlID); !isId {
		return nil, nil
	}

	base, err := unwrap(typ)
	if err != nil {
		return nil, err
	}

	switch base.Kind() {
	case reflect.Uint64, reflect.Uint, reflect.Uint32,
		reflect.Int64, reflect.Int, reflect.Int32,
		reflect.String:
		return graphql.ID, nil
	}
	return nil, nil
}
