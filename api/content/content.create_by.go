package content

import (
	"context"
	"github.com/graphql-go/graphql"
	"github.com/ichaly/go-next/lib/cms"
	"github.com/ichaly/go-next/lib/sys"
	"github.com/ichaly/go-next/pkg/base"
	"github.com/ichaly/go-next/pkg/gql"
	"gorm.io/gorm"
)

type createdBy struct {
	gql.SchemaMeta[cms.Content, *sys.User] `name:"createdBy" description:"作者信息"`

	db     *gorm.DB
	loader *gql.Loader[base.ID, *sys.User]
}

func NewContentCreatedBy(db *gorm.DB) gql.Schema {
	my := &createdBy{db: db}
	my.loader = gql.NewBatchedLoader(my.batchUsers)
	return my
}

func (my *createdBy) Resolve(p graphql.ResolveParams) (interface{}, error) {
	return base.QueryResolver[*cms.Content](p, my.db)
}

func (my *createdBy) batchUsers(ctx context.Context, keys []base.ID) []*gql.Result[*sys.User] {
	//从数据库获取数据
	var res []*sys.User
	err := my.db.WithContext(ctx).Model(&sys.User{}).Where("id in ?", keys).Find(&res).Error
	//分组数据
	values := make(map[base.ID]*sys.User)
	for _, v := range res {
		values[(*v).ID] = v
	}
	//填充结果集
	results := make([]*gql.Result[*sys.User], len(keys))
	for i, k := range keys {
		results[i] = &gql.Result[*sys.User]{
			Error: err, Data: values[k],
		}
	}
	return results
}
