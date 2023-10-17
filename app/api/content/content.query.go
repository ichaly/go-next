package content

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/go-next/app/cms"
	"github.com/ichaly/go-next/pkg/base"
	"github.com/ichaly/go-next/pkg/core"
	"gorm.io/gorm"
)

var Content = &cms.Content{}

type query struct {
	db *gorm.DB
}

func NewContentQuery(db *gorm.DB) core.Schema {
	return &query{db: db}
}

func (*query) Name() string {
	return "contents"
}

func (*query) Description() string {
	return "内容列表"
}

func (*query) Host() interface{} {
	return core.Query
}

func (*query) Type() interface{} {
	return []*cms.Content{}
}

func (my *query) Resolve(p graphql.ResolveParams) (interface{}, error) {
	return base.QueryResolver[*cms.Content](p, my.db)
}
