package content

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/go-next/app/cms"
	"github.com/ichaly/go-next/pkg/base"
	"github.com/ichaly/go-next/pkg/core"
	"gorm.io/gorm"
)

type mutation struct {
	db       *gorm.DB
	validate *base.Validate
}

func NewContentMutation(d *gorm.DB, v *base.Validate) core.Schema {
	return &mutation{db: d, validate: v}
}

func (*mutation) Name() string {
	return "contents"
}

func (*mutation) Description() string {
	return "内容管理"
}

func (*mutation) Host() interface{} {
	return core.Mutation
}

func (*mutation) Type() interface{} {
	return Content
}

func (my *mutation) Resolve(p graphql.ResolveParams) (interface{}, error) {
	return base.MutationResolver[*cms.Content](p, my.db, my.validate)
}
