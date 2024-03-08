package content

import (
	"context"
	"github.com/graphql-go/graphql"
	"github.com/ichaly/go-next/lib/base"
	"github.com/ichaly/go-next/lib/gql"
	dl "github.com/ichaly/go-next/lib/gql/dataloader"
	"github.com/ichaly/go-next/pkg/cms"
	"github.com/ichaly/go-next/pkg/sys"
	"gorm.io/gorm"
)

type creator struct {
	gql.SchemaMeta[cms.Content, *sys.User] `name:"createdBy" description:"作者信息"`

	db     *gorm.DB
	loader *dl.Loader[*base.Id, *sys.User]
}

func NewContentCreatedBy(db *gorm.DB) gql.Schema {
	my := &creator{db: db}
	my.loader = dl.NewBatchedLoader(my.batchUsers)
	return my
}

func (my *creator) Resolve(p graphql.ResolveParams) (interface{}, error) {
	uid := p.Source.(*cms.Content).CreatedBy
	thunk := my.loader.Load(p.Context, uid)
	return func() (interface{}, error) {
		return thunk()
	}, nil
}

func (my *creator) batchUsers(ctx context.Context, keys []*base.Id) []*dl.Result[*sys.User] {
	//从数据库获取数据
	var res []*sys.User
	err := my.db.WithContext(ctx).Model(&sys.User{}).Where("id in ?", keys).Find(&res).Error
	//分组数据
	values := make(map[base.Id]*sys.User)
	for _, v := range res {
		values[v.Id] = v
	}
	//填充结果集
	results := make([]*dl.Result[*sys.User], len(keys))
	for i, k := range keys {
		r := &dl.Result[*sys.User]{Error: err}
		if k != nil {
			r.Data = values[*k]
		}
		results[i] = r
	}
	return results
}
