package user

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/go-next/app/sys"
	"github.com/ichaly/go-next/pkg/base"
	"github.com/ichaly/go-next/pkg/core"
	"github.com/ichaly/go-next/pkg/util"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var User = &sys.User{}

type query struct {
	db *gorm.DB
}

func NewUserQuery(db *gorm.DB) core.Schema {
	return &query{db: db}
}

func (*query) Name() string {
	return "users"
}

func (*query) Description() string {
	return "用户列表"
}

func (*query) Host() interface{} {
	return core.Query
}

func (my *query) Type() interface{} {
	return []*sys.User{}
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
