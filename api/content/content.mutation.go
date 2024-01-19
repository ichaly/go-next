package content

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/go-next/lib/cms"
	"github.com/ichaly/go-next/pkg/base"
	"github.com/ichaly/go-next/pkg/gql"
	"gorm.io/gorm"
)

type mutation struct {
	gql.SchemaMeta[gql.Mutation, *cms.Content] `name:"contents" description:"内容管理"`

	db       *gorm.DB
	validate *base.Validate
}

func NewContentMutation(d *gorm.DB, v *base.Validate) gql.Schema {
	return &mutation{db: d, validate: v}
}

func (my *mutation) Resolve(p graphql.ResolveParams) (interface{}, error) {
	return base.MutationResolver[*cms.Content](p, my.db, my.validate)
}
