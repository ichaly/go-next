package gql

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"reflect"
)

type Engine struct {
	options Options
	types   map[string]graphql.Type
}

func NewEngine(opts ...EngineOption) *Engine {
	return &Engine{types: map[string]graphql.Type{}, options: newOptions(opts...)}
}

func (my *Engine) Register(node Schema) error {
	if node == nil {
		return fmt.Errorf("node can't be nil")
	}

	nodeType, err := unwrap(reflect.TypeOf(node))
	if err != nil {
		return err
	}

	if nodeType.Kind() != reflect.Struct {
		return fmt.Errorf("reflect: NumField of non-struct type " + nodeType.String())
	}

	field, ok := nodeType.FieldByName(SCHEMA_META)
	if !ok {
		return fmt.Errorf("field 'Schema Meta' not found")
	}

	name, description := field.Tag.Get(TAG_NAME), field.Tag.Get(TAG_DESCRIPTION)
	if name == "" {
		name = nodeType.Name()
	}
	parentField, _ := nodeType.FieldByName(PARENT_TYPE)
	resultField, _ := nodeType.FieldByName(RESULT_TYPE)

	parentType, err := parseType(parentField.Type, "parent", my.asObject)
	if err != nil {
		return err
	}
	resultType, err := parseType(resultField.Type, "result",
		my.asBuiltinScalar, my.asCustomScalar, my.asEnum, my.asId, my.asObject,
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
			"sort":  {Type: my.types[resultType.Name()+SUFFIX_SORT_INPUT]},
			"where": {Type: my.types[resultType.Name()+SUFFIX_WHERE_INPUT]},
		}
	}
	if parent.Name() == MUTATION {
		args["data"] = &graphql.ArgumentConfig{Type: my.types[resultType.Name()+SUFFIX_DATA_INPUT]}
		args["delete"] = &graphql.ArgumentConfig{Type: graphql.Boolean}
	}
	parent.AddFieldConfig(name, &graphql.Field{
		Type: wrapType(resultField.Type, resultType), Args: args, Resolve: node.Resolve, Description: description,
	})
	return nil
}

func (my *Engine) Schema() (graphql.Schema, error) {
	config := graphql.SchemaConfig{}
	if q := my.checkSchema(QUERY); q != nil {
		config.Query = q
	}
	if m := my.checkSchema(MUTATION); m != nil {
		config.Mutation = m
	}
	if s := my.checkSchema(SUBSCRIPTION); s != nil {
		config.Subscription = s
	}
	return graphql.NewSchema(config)
}

func (my *Engine) checkSchema(name string) *graphql.Object {
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
