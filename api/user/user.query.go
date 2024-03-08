package user

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/go-next/lib/base"
	"github.com/ichaly/go-next/lib/gql"
	"github.com/ichaly/go-next/lib/util"
	"github.com/ichaly/go-next/pkg/sys"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type query struct {
	gql.SchemaMeta[gql.Query, []*sys.User] `name:"users" description:"用户列表"`

	db *gorm.DB
}

func NewUserQuery(db *gorm.DB) gql.Schema {
	return &query{db: db}
}

func (my *query) Resolve(p graphql.ResolveParams) (interface{}, error) {
	password := util.EraseMap(p.Args, "where.password")
	res, err := base.QueryResolver[*sys.User](p, my.db)
	if err != nil {
		return nil, err
	}
	if password == nil {
		return res, nil
	}
	if users, ok := res.([]*sys.User); ok {
		for _, u := range users {
			err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password.(map[string]interface{})["eq"].(string)))
			if err == nil {
				return []*sys.User{u}, nil
			}
		}
	}
	return nil, nil
}
