package sys

import (
	"fmt"
	"github.com/ichaly/go-next/lib/base"
	"gorm.io/gorm"
)

type RoleRule struct {
	RoleId       base.Id `gorm:"primaryKey;comment:角色ID"`
	RuleId       base.Id `gorm:"primaryKey;comment:菜单ID"`
	base.General `mapstructure:",squash"`
}

func (*RoleRule) TableName() string {
	return "sys_role_rule"
}

// AfterSave 添加关联后，创建casbin的权限与角色的关联
func (my *RoleRule) AfterSave(tx *gorm.DB) (err error) {
	var rule *Rule
	if e := tx.Where("id = ?", my.RuleId).First(&rule).Error; e != nil {
		err = fmt.Errorf("关联权限和角色获取异常: <%s>", e.Error())
	}
	if _, e := roleService.AddRolePolicy(my.RoleId, *rule.Code, rule.Action); e != nil {
		err = fmt.Errorf("关联权限和角色到casbin异常: <%s>", e.Error())
	}
	return
}
