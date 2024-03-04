package role

import (
	"github.com/graphql-go/graphql"
	"github.com/ichaly/go-next/lib/sys"
	"github.com/ichaly/go-next/pkg/base"
	"github.com/ichaly/go-next/pkg/gql"
	"github.com/ichaly/go-next/pkg/util"
	"gorm.io/gorm"
)

type mutation struct {
	gql.SchemaMeta[gql.Mutation, *sys.Role] `name:"roles" description:"角色管理"`

	db       *gorm.DB
	validate *base.Validate
}

func NewRoleMutation(d *gorm.DB, v *base.Validate) gql.Schema {
	return &mutation{db: d, validate: v}
}

func (my *mutation) Resolve(p graphql.ResolveParams) (interface{}, error) {
	//处理参数
	data := p.Args["data"].(map[string]interface{})
	ruleIds := data["ruleIds"].([]interface{})
	delete(data, "ruleIds")

	//开启事务
	tx := my.db.Begin()
	//创建或修改角色
	tmp, err := base.MutationResolver[*sys.Role](p, tx, my.validate)
	if err != nil {
		return nil, err
	}

	//删除原来的权限关系
	role := tmp.(*sys.Role)
	tx.Delete(&sys.RoleRule{}, "role_id = ?", role.Id)

	//构建新的角色权限关系
	var rules []sys.RoleRule
	for _, v := range ruleIds {
		ruleId := base.Id(util.ParseLong(v.(string)))
		rules = append(rules, sys.RoleRule{RoleId: role.Id, RuleId: ruleId})
	}

	//保存新关系
	err = tx.Create(&rules).Error
	if err != nil {
		return nil, err
	}

	//提交事务
	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}
	return tmp, nil
}
