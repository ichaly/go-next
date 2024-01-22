package gql

import (
	"github.com/graphql-go/graphql"
)

func (my *Engine) buildExpressionInput(t graphql.Type) graphql.Type {
	var name string
	var list []input
	var isList bool

	list = append(list, expNull...)
	if val, ok := t.(*graphql.NonNull); ok {
		t = val.OfType
	}
	if val, ok := t.(*graphql.List); ok {
		t = val.OfType
		isList = true
	}
	if val, ok := t.(*graphql.NonNull); ok {
		t = val.OfType
	}
	//if val, ok := t.(*graphql.Enum); ok {
	//	t = val
	//	isEnum = true
	//}
	if val, ok := t.(*graphql.Object); ok {
		return my.buildWhereInput(val)
	}

	if isList {
		list = append(list, expList...)
	} else {
		list = append(list, expBase...)
	}

	if isList {
		name = t.Name() + SUFFIX_EXPR_LIST
	} else {
		name = t.Name() + SUFFIX_EXPRESSION
	}
	typ, ok := my.types[name]
	if ok {
		return typ
	}

	fields := graphql.InputObjectConfigFieldMap{}
	for _, e := range list {
		if e.Type == nil {
			e.Type = t
		}
		fields[e.Name] = &graphql.InputObjectFieldConfig{Type: e.Type, Description: e.Desc}
	}
	input := graphql.NewInputObject(graphql.InputObjectConfig{Name: name, Fields: fields})
	my.types[name] = input

	return input
}

func (my *Engine) buildDataInput(object *graphql.Object) graphql.Type {
	name := object.Name() + SUFFIX_DATA_INPUT
	val, ok := my.types[name]
	if ok {
		return val
	}

	fields := graphql.InputObjectConfigFieldMap{}
	for k, f := range object.Fields() {
		fields[k] = &graphql.InputObjectFieldConfig{Type: f.Type, Description: f.Description}
	}
	input := graphql.NewInputObject(graphql.InputObjectConfig{
		Name: name, Fields: fields,
	})

	my.types[name] = input
	return input
}

func (my *Engine) buildSortInput(object *graphql.Object) graphql.Type {
	name := object.Name() + SUFFIX_SORT_INPUT
	val, ok := my.types[name]
	if ok {
		return val
	}

	fields := graphql.InputObjectConfigFieldMap{}
	for k := range object.Fields() {
		fields[k] = &graphql.InputObjectFieldConfig{Type: SortDirection}
	}
	input := graphql.NewInputObject(graphql.InputObjectConfig{
		Name: name, Fields: fields,
	})

	my.types[name] = input
	return input
}

func (my *Engine) buildWhereInput(object *graphql.Object) graphql.Type {
	name := object.Name() + SUFFIX_WHERE_INPUT
	val, ok := my.types[name]
	if ok {
		return val
	}
	fields := graphql.InputObjectConfigFieldMap{}
	for k, v := range object.Fields() {
		typ := my.buildExpressionInput(v.Type)
		if typ != nil {
			fields[k] = &graphql.InputObjectFieldConfig{Type: typ}
		}
	}
	input := graphql.NewInputObject(graphql.InputObjectConfig{
		Name: name, Fields: fields,
	})

	fields["or"] = &graphql.InputObjectFieldConfig{Type: input}
	fields["and"] = &graphql.InputObjectFieldConfig{Type: input}
	fields["not"] = &graphql.InputObjectFieldConfig{Type: input}

	my.types[name] = input
	return input
}
