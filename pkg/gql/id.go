package gql

import (
	"github.com/graphql-go/graphql"
	"reflect"
)

type GqlID interface {
	ID()
}

func (my *Engine) asId(typ reflect.Type) (graphql.Type, error) {
	base, err := unwrap(typ)
	if err != nil {
		return nil, err
	}

	if _, isId := isImplements[GqlID](base); isId {
		return ID, nil
	}

	switch base.Kind() {
	case reflect.Uint64, reflect.Uint, reflect.Uint32,
		reflect.Int64, reflect.Int, reflect.Int32,
		reflect.String:
		return ID, nil
	default:
		return nil, nil
	}
}
