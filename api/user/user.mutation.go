package user

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/go-next/lib/base"
	"github.com/ichaly/go-next/lib/gql"
	"github.com/ichaly/go-next/pkg/sys"
	"gorm.io/gorm"
)

type mutation struct {
	gql.SchemaMeta[gql.Mutation, *sys.User] `name:"users" description:"用户管理"`

	db       *gorm.DB
	validate *base.Validate
}

func NewUserMutation(d *gorm.DB, v *base.Validate) gql.Schema {
	return &mutation{db: d, validate: v}
}

func (my *mutation) Resolve(p graphql.ResolveParams) (interface{}, error) {
	return base.MutationResolver[*sys.User](p, my.db, my.validate)
}
