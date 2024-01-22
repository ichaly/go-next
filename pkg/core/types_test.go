package core

import (
	"reflect"
	"testing"
)

type foo struct{}

type boo struct {
	SchemaMeta[*foo, []int] `name:"users" description:"用户列表"`
}

func TestSchemaMeta(t *testing.T) {
	b := boo{}
	typ := reflect.TypeOf(b)
	t.Log(reflect.TypeOf(b.parentType).Elem().Name())
	t.Log(reflect.TypeOf(b.resultType).Elem().Name())
	if field, ok := typ.FieldByName("SchemaMeta"); ok {
		t.Log(field.Tag.Get("name"))
		t.Log(field.Tag.Get("description"))
	}
}
