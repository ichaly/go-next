package rule

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/go-next/lib/sys"
	"github.com/ichaly/go-next/pkg/base"
	"github.com/ichaly/go-next/pkg/gql"
	"gorm.io/gorm"
)

type mutation struct {
	gql.SchemaMeta[gql.Mutation, *sys.Rule] `name:"rules" description:"权限管理"`

	db       *gorm.DB
	validate *base.Validate
}

func NewRuleMutation(d *gorm.DB, v *base.Validate) gql.Schema {
	return &mutation{db: d, validate: v}
}

func (my *mutation) Resolve(p graphql.ResolveParams) (interface{}, error) {
	return base.MutationResolver[*sys.Rule](p, my.db, my.validate)
}
