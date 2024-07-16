package rule

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/go-next/lib/base"
	"github.com/ichaly/go-next/lib/gql"
	"github.com/ichaly/go-next/pkg/sys"
	"gorm.io/gorm"
)

type query struct {
	gql.SchemaMeta[gql.Query, []*sys.Rule] `name:"rules" description:"权限列表"`

	db *gorm.DB
}

func NewRuleQuery(db *gorm.DB) gql.Schema {
	return &query{db: db}
}

func (my *query) Resolve(p graphql.ResolveParams) (interface{}, error) {
	return base.QueryResolver[*sys.Rule](p, my.db)
}