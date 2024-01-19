package user

import (
	"context"
	"github.com/graphql-go/graphql"
	"github.com/ichaly/go-next/lib/cms"
	"github.com/ichaly/go-next/lib/sys"
	"github.com/ichaly/go-next/pkg/gql"
	"gorm.io/gorm"
)

type contents struct {
	db     *gorm.DB
	loader *gql.Loader[uint64, []*cms.Content]
}

func NewUserContents(db *gorm.DB) gql.Schema {
	my := &contents{db: db}
	my.loader = gql.NewBatchedLoader(my.batchContents)
	return my
}

func (*contents) Name() string {
	return "contents"
}

func (*contents) Description() string {
	return "用户作品"
}

func (*contents) Host() interface{} {
	return User
}

func (*contents) Type() interface{} {
	return []*cms.Content{}
}

func (my *contents) Resolve(p graphql.ResolveParams) (interface{}, error) {
	uid := uint64(p.Source.(*sys.User).ID)
	thunk := my.loader.Load(p.Context, uid)
	return func() (interface{}, error) {
		return thunk()
	}, nil
}

func (my *contents) batchContents(ctx context.Context, keys []uint64) []*gql.Result[[]*cms.Content] {
	//从数据库获取数据
	var res []*cms.Content
	err := my.db.WithContext(ctx).Model(&cms.Content{}).Where("created_by in ?", keys).Find(&res).Error
	//分组数据
	values := make(map[uint64][]*cms.Content)
	for _, c := range res {
		values[*c.CreatedBy] = append(values[*c.CreatedBy], c)
	}
	//填充结果集
	results := make([]*gql.Result[[]*cms.Content], len(keys))
	for i, k := range keys {
		results[i] = &gql.Result[[]*cms.Content]{
			Error: err, Data: values[k],
		}
	}
	return results
}
