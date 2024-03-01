package sys

import (
	"fmt"
	"github.com/ichaly/go-next/pkg/base"
	"gorm.io/gorm"
)

type RoleRule struct {
	base.Primary `mapstructure:",squash"`
	RoleId       base.Id `gorm:"comment:角色ID"`
	RuleId       base.Id `gorm:"comment:菜单ID"`
}

func (*RoleRule) TableName() string {
	return "sys_role_rule"
}

// AfterCreate 添加关联后，创建casbin的权限与角色的关联
func (my *RoleRule) AfterCreate(tx *gorm.DB) (err error) {
	var rule *Rule
	if e := tx.Where("mid = ?", my.RuleId).First(rule).Error; e != nil {
		err = fmt.Errorf("关联权限和角色获取异常: <%s>", e.Error())
	}
	if _, e := roleService.AddRolePolicy(my.RoleId, rule.Code, rule.Action); e != nil {
		err = fmt.Errorf("关联权限和角色到casbin异常: <%s>", e.Error())
	}
	return
}
