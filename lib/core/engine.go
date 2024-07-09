package core

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/go-next/lib/core/internal"
	"github.com/spf13/viper"
)

type Engine struct {
	meta  *Metadata
	types map[string]graphql.Type
}

func NewEngine(m *Metadata, v *viper.Viper) (*Engine, error) {
	my := &Engine{meta: m, types: map[string]graphql.Type{}}
	//启动引擎
	my.start()
	return my, nil
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
