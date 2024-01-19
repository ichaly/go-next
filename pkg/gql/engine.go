package gql

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"reflect"
)

const (
	QUERY        = "Query"
	MUTATION     = "Mutation"
	SUBSCRIPTION = "Subscription"
)

type Engine struct {
	types map[string]graphql.Type
}

func NewEngine() *Engine {
	return &Engine{types: map[string]graphql.Type{}}
}

func (my *Engine) Register(node Schema) error {
	if node == nil {
		return fmt.Errorf("node can't be nil")
	}

	nodeType := reflect.TypeOf(node)
	if nodeType.Kind() == reflect.Ptr {
		nodeType = nodeType.Elem()
	}
	field, ok := nodeType.FieldByName("SchemaMeta")
	if !ok {
		return fmt.Errorf("field 'Schema Meta' not found")
	}

	name := field.Tag.Get("name")
	if name == "" {
		name = nodeType.Name()
	}
	description := field.Tag.Get("description")
	parentField, _ := nodeType.FieldByName("parentType")
	resultField, _ := nodeType.FieldByName("resultType")

	resultType, err := parseType(resultField.Type, "result",
		my.asBuiltinScalar, my.asCustomScalar, my.asId, my.asEnum, my.asObject,
	)
	if err != nil {
		return err
	}
	parentType, err := parseType(parentField.Type, "parent",
		my.asBuiltinScalar, my.asCustomScalar, my.asId, my.asEnum, my.asObject,
	)
	if err != nil {
		return err
	}

	parent, ok := parentType.(*graphql.Object)
	if !ok {
		return fmt.Errorf("invalid parent %s", parentField.Type)
	}

	var args graphql.FieldConfigArgument
	if _, ok := resultType.(*graphql.Object); ok {
		args = graphql.FieldConfigArgument{
			"size":  {Type: graphql.Int},
			"page":  {Type: graphql.Int},
			"sort":  {Type: my.types[resultType.Name()+"SortInput"]},
			"where": {Type: my.types[resultType.Name()+"WhereInput"]},
		}
	}
	if parent.Name() == MUTATION {
		args["data"] = &graphql.ArgumentConfig{Type: my.types[resultType.Name()+"DataInput"]}
		args["delete"] = &graphql.ArgumentConfig{Type: graphql.Boolean}
	}
	parent.AddFieldConfig(name, &graphql.Field{
		Type: wrapType(resultField.Type, resultType), Args: args, Resolve: node.Resolve, Description: description,
	})
	return nil
}

func (my *Engine) Schema() (graphql.Schema, error) {
	config := graphql.SchemaConfig{}
	if q := my.checkObject(QUERY); q != nil {
		config.Query = q
	}
	if m := my.checkObject(MUTATION); m != nil {
		config.Mutation = m
	}
	if s := my.checkObject(SUBSCRIPTION); s != nil {
		config.Subscription = s
	}
	return graphql.NewSchema(config)
}

func (my *Engine) checkObject(name string) *graphql.Object {
	val, ok := my.types[name]
	if !ok {
		return nil
	}
	obj, ok := val.(*graphql.Object)
	if !ok {
		return nil
	}
	if len(obj.Fields()) <= 0 {
		return nil
	}
	return obj
}
