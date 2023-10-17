package user

import (
	"context"
	"github.com/graphql-go/graphql"
	"github.com/ichaly/go-next/app/cms"
	"github.com/ichaly/go-next/app/sys"
	"github.com/ichaly/go-next/pkg/core"
	"gorm.io/gorm"
)

type contents struct {
	db     *gorm.DB
	loader *core.Loader[uint64, []*cms.Content]
}

func NewUserContents(db *gorm.DB) core.Schema {
	my := &contents{db: db}
	my.loader = core.NewBatchedLoader(my.batchContents)
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

func (my *contents) batchContents(ctx context.Context, keys []uint64) []*core.Result[[]*cms.Content] {
	var res []*cms.Content
	err := my.db.WithContext(ctx).Model(&cms.Content{}).Where("created_by in ?", keys).Find(&res).Error
	values := make(map[uint64][]*cms.Content)
	for _, c := range res {
		values[*c.CreatedBy] = append(values[*c.CreatedBy], c)
	}
	results := make([]*core.Result[[]*cms.Content], len(keys))
	for i, k := range keys {
		r := &core.Result[[]*cms.Content]{
			Error: err,
		}
		if v, ok := values[k]; ok {
			r.Data = v
		}
		results[i] = r
	}
	return results
}
