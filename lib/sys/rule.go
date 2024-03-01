package sys

import (
	"fmt"
	"github.com/ichaly/go-next/pkg/base"
	"gorm.io/gorm"
)

const (
	RoleRoot  = "root"
	RoleGuest = "guest"
)

const (
	ActionAll   = "*"
	ActionRead  = "read"
	ActionWrite = "write"
)

type Rule struct {
	Pid         *base.Id `gorm:"comment:父级ID" json:",omitempty"`
	Kind        RuleKind `gorm:"comment:权限类型"`
	Code        string   `gorm:"index;size:200;comment:唯一标识"`
	Icon        string   `gorm:"size:100;comment:图标"`
	Title       string   `gorm:"size:200;comment:标题"`
	Action      string   `gorm:"size:100;comment:动作"`
	Weight      int8     `gorm:"comment:权重"`
	Hidden      bool     `gorm:"comment:是否隐藏"`
	Default     bool     `gorm:"comment:是否默认"`
	External    bool     `gorm:"comment:是否外链"`
	base.Entity `mapstructure:",squash"`
}

func (*Rule) TableName() string {
	return "sys_rule"
}

// BeforeDelete 删除用户前清除用户与角色的关联信息
func (my *Rule) BeforeDelete(tx *gorm.DB) (err error) {
	// 清除casbin权限与角色关联
	if _, e := roleService.DeleteRule(my.Code); e != nil {
		err = fmt.Errorf("清除casbin角色权限异常: <%s>", e.Error())
	}
	if e := tx.Where("mid = ?", my.Id).Delete(&RoleRule{}).Error; e != nil {
		err = fmt.Errorf("删除角色权限关联异常: <%s>", e.Error())
	}
	return
}
