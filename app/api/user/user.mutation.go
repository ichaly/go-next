package user

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/go-next/app/sys"
	"github.com/ichaly/go-next/pkg/base"
	"github.com/ichaly/go-next/pkg/gql"
	"gorm.io/gorm"
)

type mutation struct {
	db       *gorm.DB
	validate *base.Validate
}

func NewUserMutation(d *gorm.DB, v *base.Validate) gql.Schema {
	return &mutation{db: d, validate: v}
}

func (*mutation) Name() string {
	return "users"
}

func (*mutation) Description() string {
	return "用户管理"
}

func (*mutation) Host() interface{} {
	return gql.Mutation
}

func (*mutation) Type() interface{} {
	return User
}

func (my *mutation) Resolve(p graphql.ResolveParams) (interface{}, error) {
	return base.MutationResolver[*sys.User](p, my.db, my.validate)
}
