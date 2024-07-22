package core

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/iancoleman/strcase"
	"github.com/ichaly/go-next/lib/core/internal"
)

const (
	QUERY        = "Query"
	MUTATION     = "Mutation"
	SUBSCRIPTION = "Subscription"
)

type Option func(*graphql.ObjectConfig)

func WithHost(host string) Option {
	return func(c *graphql.ObjectConfig) {
		c.Name = host
	}
}

func WithDescription(description string) Option {
	return func(c *graphql.ObjectConfig) {
		c.Description = description
	}
}

type Engine struct {
	meta  *Metadata
	reso  *Resolver
	types map[string]*graphql.Object
}

func NewEngine(m *Metadata, r *Resolver) (*Engine, error) {
	my := &Engine{meta: m, reso: r, types: make(map[string]*graphql.Object)}

	//启动引擎
	for _, t := range my.meta.Nodes {
		//注册列
		for _, c := range t.Columns {
			_ = my.Register(
				&graphql.Field{
					Name:        c.Name,
					Type:        my.parseType(c),
					Resolve:     my.reso.Resolve,
					Description: c.Description,
				},
				WithHost(t.Name),
				WithDescription(t.Description),
			)
		}

		//注册表
		_ = my.Register(
			&graphql.Field{
				Name:        strcase.ToLowerCamel(t.Name),
				Type:        my.types[t.Name],
				Description: t.Description,
			},
		)
	}

	return my, nil
}

func (my *Engine) Register(f *graphql.Field, options ...Option) error {
	if f == nil {
		return errors.New("field is nil")
	}

	cfg := &graphql.ObjectConfig{Name: QUERY, Fields: graphql.Fields{}}
	for _, option := range options {
		option(cfg)
	}

	host, ok := my.types[cfg.Name]
	if !ok {
		host = graphql.NewObject(*cfg)
		my.types[cfg.Name] = host
	}

	host.AddFieldConfig(f.Name, f)
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
	obj, ok := my.types[name]
	if !ok {
		return nil
	}
	if len(obj.Fields()) <= 0 {
		return nil
	}
	return obj
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
