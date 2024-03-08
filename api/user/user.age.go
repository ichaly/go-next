package user

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/go-next/lib/gql"
	"github.com/ichaly/go-next/pkg/sys"
	"time"
)

type age struct {
	gql.SchemaMeta[sys.User, int] `name:"age" description:"年龄"`
}

func NewUserAge() gql.Schema {
	return &age{}
}

func (my *age) Resolve(p graphql.ResolveParams) (interface{}, error) {
	user := p.Source.(*sys.User)
	if user.Birthday == nil || user.Birthday.IsZero() {
		return 0, nil
	}
	return time.Now().Year() - user.Birthday.Year(), nil
}
