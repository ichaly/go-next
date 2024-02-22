package sys

import (
	"fmt"
	"github.com/ichaly/go-next/pkg/base"
	"gorm.io/gorm"
)

type Role struct {
	Code   string `gorm:"index:,unique;size:200;comment:编码"`
	Name   string `gorm:"size:200;comment:名称"`
	Scope  string `gorm:"size:200;comment:数据权限"`
	Weight string `gorm:"comment:权重"`
	base.Entity
}

func (*Role) TableName() string {
	return "sys_role"
}

// BeforeDelete 删除角色前，清除角色与用户的关联
func (my *Role) BeforeDelete(tx *gorm.DB) (err error) {
	// 清除casbin用户与角色关联
	if _, e := roleService.DeleteRole(my.Id); e != nil {
		err = fmt.Errorf("清除casbin角色权限异常: <%s>", e.Error())
	}
	// 清除数据库中用户与角色的关联
	if e := tx.Where("rid = ?", my.Id).Delete(&RoleUser{}).Error; e != nil {
		err = fmt.Errorf("删除用户角色关联异常: <%s>", e.Error())
	}
	return
}
