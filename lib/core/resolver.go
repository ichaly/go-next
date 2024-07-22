package core

import (
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

type Resolver struct {
	db   *gorm.DB
	meta *Metadata
}

func NewResolver(m *Metadata, d *gorm.DB) (*Resolver, error) {
	return &Resolver{db: d, meta: m}, nil
}

func (my *Resolver) Resolve(p graphql.ResolveParams) (interface{}, error) {
	return nil, nil
}
