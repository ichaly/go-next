package user

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

type contents struct {
	gql.SchemaMeta[sys.User, []*cms.Content] `name:"contents" description:"用户作品"`

	db     *gorm.DB
	loader *dl.Loader[base.Id, []*cms.Content]
}

func NewUserContents(db *gorm.DB) gql.Schema {
	my := &contents{db: db}
	my.loader = dl.NewBatchedLoader(my.batchContents)
	return my
}

func (my *contents) Resolve(p graphql.ResolveParams) (interface{}, error) {
	uid := p.Source.(*sys.User).Id
	thunk := my.loader.Load(p.Context, uid)
	return func() (interface{}, error) {
		return thunk()
	}, nil
}

func (my *contents) batchContents(ctx context.Context, keys []base.Id) []*dl.Result[[]*cms.Content] {
	//从数据库获取数据
	var res []*cms.Content
	err := my.db.WithContext(ctx).Model(&cms.Content{}).Where("created_by in ?", keys).Find(&res).Error
	//分组数据
	values := make(map[*base.Id][]*cms.Content)
	for _, c := range res {
		values[c.CreatedBy] = append(values[c.CreatedBy], c)
	}
	//填充结果集
	results := make([]*dl.Result[[]*cms.Content], len(keys))
	for i, k := range keys {
		results[i] = &dl.Result[[]*cms.Content]{
			Error: err, Data: values[&k],
		}
	}
	return results
}
