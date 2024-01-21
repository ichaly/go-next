package core

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"reflect"
)

func isImplements[V any](t reflect.Type) (V, bool) {
	v, ok := reflect.New(t).Interface().(V)
	return v, ok
}

func isNative(p reflect.Type) bool {
	switch p.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Bool,
		reflect.String,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.Map,
		reflect.Struct,
		reflect.Interface:
		return true
	default:
		return false
	}
}

func unwrap(p reflect.Type) (reflect.Type, error) {
	switch p.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Array:
		return unwrap(p.Elem())
	default:
		if !isNative(p) {
			return nil, fmt.Errorf("unsupported type('%s') to unwrap", p.String())
		}
		return p, nil
	}
}

func wrapType(p reflect.Type, t graphql.Type) graphql.Type {
	var isPtr bool
	if p.Kind() == reflect.Ptr {
		isPtr = true
		p = p.Elem()
	}
	if p.Kind() == reflect.Slice || p.Kind() == reflect.Array {
		t = graphql.NewList(wrapType(p.Elem(), t))
	}
	if !isPtr {
		t = graphql.NewNonNull(t)
	}
	return t
}

func parseType(typ reflect.Type, errString string, parsers ...typeParser) (graphql.Type, error) {
	for _, parse := range parsers {
		res, err := parse(typ)
		if err != nil {
			return nil, err
		}
		if res == nil {
			continue
		}
		return res, nil
	}
	return nil, fmt.Errorf("unsupported type('%s') for %s", typ.String(), errString)
}
