package gql

import (
	"github.com/graphql-go/graphql"
	"reflect"
	"strings"
)

type EngineOption func(*Options)

type Options struct {
	ignoreField              []string
	defaultResolver          func(p graphql.ResolveParams) (interface{}, error)
	fieldDescriptionProvider func(field *reflect.StructField) string
}

func newOptions(opts ...EngineOption) Options {
	opt := Options{
		fieldDescriptionProvider: func(field *reflect.StructField) string {
			tag := field.Tag.Get("gorm")
			tags := strings.Split(tag, ";")
			for _, t := range tags {
				if strings.HasPrefix(t, "comment:") {
					return t[8:]
				}
			}
			return ""
		},
		ignoreField:     []string{"password", "secret"},
		defaultResolver: defaultResolver,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

func WithFieldDescriptionProvider(f func(*reflect.StructField) string) EngineOption {
	return func(o *Options) {
		if f != nil {
			o.fieldDescriptionProvider = f
		}
	}
}

func WithDefaultResolver(f func(p graphql.ResolveParams) (interface{}, error)) EngineOption {
	return func(o *Options) {
		if f != nil {
			o.defaultResolver = f
		}
	}
}
func WithIgnoreField(a []string) EngineOption {
	return func(o *Options) {
		if len(a) > 0 {
			o.ignoreField = a
		}
	}
}
