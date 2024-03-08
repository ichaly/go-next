package content

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/go-next/lib/base"
	"github.com/ichaly/go-next/lib/gql"
	"github.com/ichaly/go-next/pkg/cms"
	"gorm.io/gorm"
)

type query struct {
	gql.SchemaMeta[gql.Query, []*cms.Content] `name:"contents" description:"内容列表"`

	db *gorm.DB
}

func NewContentQuery(db *gorm.DB) gql.Schema {
	return &query{db: db}
}

func (my *query) Resolve(p graphql.ResolveParams) (interface{}, error) {
	return base.QueryResolver[*cms.Content](p, my.db)
}
