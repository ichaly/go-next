package core

import (
	"reflect"
	"strings"
)

type EngineOption func(*Options)

type Options struct {
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
