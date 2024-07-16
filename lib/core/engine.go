package core

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/ichaly/go-next/lib/core/internal"
)

const (
	QUERY        = "Query"
	MUTATION     = "Mutation"
	SUBSCRIPTION = "Subscription"
)

type QueryField interface {
	Field() *graphql.Field
}

type NamedField interface {
	QueryField
	Name() string
}

type Engine struct {
	meta  *Metadata
	types map[string]graphql.Type
}

func NewEngine(m *Metadata) (*Engine, error) {
	my := &Engine{meta: m, types: map[string]graphql.Type{}}
	//启动引擎
	my.start()
	return my, nil
}

func (my *Engine) Register(s QueryField) error {
	var name string
	if value, ok := s.(NamedField); !ok {
		name = value.Name()
	} else {
		name = QUERY
	}
	if s.Field() == nil {
		return errors.New("field is nil")
	}
	host, ok := my.types[name]
	if !ok {
		if name == QUERY {
			host = graphql.NewObject(graphql.ObjectConfig{Name: QUERY})
			my.types[name] = host
		} else {
			return errors.New("host not found")
		}
	}
	node, ok := host.(*graphql.Object)
	if !ok {
		return errors.New("host is not object")
	}
	node.AddFieldConfig(s.Field().Name, s.Field())
	return nil
}

func (my *Engine) start() {
	for _, v := range my.meta.Nodes {
		fields := graphql.Fields{}
		for _, c := range v.Columns {
			fields[c.Name] = &graphql.Field{
				Type:        my.parseType(c),
				Description: c.Description,
			}
		}
		object := graphql.NewObject(graphql.ObjectConfig{
			Name: v.Name, Description: v.Description, Fields: fields,
		})
		my.types[v.Name] = object

		_ = my.Register(&internal.RootField{Object: object})
	}
}

func (my *Engine) parseType(c *Column) graphql.Type {
	if c.IsPrimary {
		return graphql.ID
	}
	switch internal.DataTypes[c.Type] {
	case graphql.Int.Name():
		return graphql.Int
	case graphql.Float.Name():
		return graphql.Float
	case graphql.Boolean.Name():
		return graphql.Boolean
	case graphql.DateTime.Name():
		return graphql.DateTime
	default:
		return graphql.String
	}
}
