package content

import (
	"context"
	"github.com/graphql-go/graphql"
	"github.com/ichaly/go-next/lib/cms"
	"github.com/ichaly/go-next/lib/sys"
	"github.com/ichaly/go-next/pkg/gql"
	"gorm.io/gorm"
)

type createdBy struct {
	gql.SchemaMeta[cms.Content, *sys.User] `name:"createdBy" description:"作者信息"`

	db     *gorm.DB
	loader *gql.Loader[*uint64, *sys.User]
}

func NewContentCreatedBy(db *gorm.DB) gql.Schema {
	my := &createdBy{db: db}
	my.loader = gql.NewBatchedLoader(my.batchUsers)
	return my
}

func (my *createdBy) Resolve(p graphql.ResolveParams) (interface{}, error) {
	uid := p.Source.(*cms.Content).CreatedBy
	thunk := my.loader.Load(p.Context, uid)
	return func() (interface{}, error) {
		return thunk()
	}, nil
}

func (my *createdBy) batchUsers(ctx context.Context, keys []*uint64) []*gql.Result[*sys.User] {
	//从数据库获取数据
	var res []*sys.User
	err := my.db.WithContext(ctx).Model(&sys.User{}).Where("id in ?", keys).Find(&res).Error
	//分组数据
	values := make(map[uint64]*sys.User)
	for _, v := range res {
		values[uint64(v.ID)] = v
	}
	//填充结果集
	results := make([]*gql.Result[*sys.User], len(keys))
	for i, k := range keys {
		r := &gql.Result[*sys.User]{Error: err}
		if k != nil {
			r.Data = values[*k]
		}
		results[i] = r
	}
	return results
}
