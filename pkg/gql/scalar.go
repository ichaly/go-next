package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"reflect"
)

type GqlScalar interface {
	GqlDescription
	Marshal() interface{}
	Unmarshal(value interface{})
}

func (my *Engine) asCustomScalar(typ reflect.Type) (graphql.Type, error) {
	if _, isScalar := reflect.New(typ).Interface().(GqlScalar); !isScalar {
		return nil, nil
	}

	scalar, err := my.buildCustomScalar(typ)
	if err != nil {
		return nil, err
	}

	return scalar, nil
}

func (my *Engine) buildCustomScalar(base reflect.Type) (*graphql.Scalar, error) {
	typ, err := unwrap(base)
	if err != nil {
		return nil, err
	}

	if val, ok := my.types[typ.Name()]; ok {
		return val.(*graphql.Scalar), nil
	}

	ptr := newPrototype(typ).(GqlScalar)

	name, desc := typ.Name(), ptr.Description()
	s := graphql.NewScalar(graphql.ScalarConfig{
		Name: name, Description: desc,
		Serialize: func(value interface{}) interface{} {
			if s, ok := value.(GqlScalar); ok {
				return s.Marshal()
			}
			return nil
		},
		ParseValue: func(value interface{}) interface{} {
			s := newPrototype(typ).(GqlScalar)
			s.Unmarshal(value)
			return s
		},
		ParseLiteral: func(value ast.Value) interface{} {
			s := newPrototype(typ).(GqlScalar)
			s.Unmarshal(value.GetValue())
			return s
		},
	})
	my.types[name] = s
	return s, nil
}
