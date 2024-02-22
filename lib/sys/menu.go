package sys

import (
	"fmt"
	"github.com/ichaly/go-next/pkg/base"
	"gorm.io/gorm"
)

type Menu struct {
	Code   string `gorm:"index:,unique;size:200;comment:编码"`
	Name   string `gorm:"size:200;comment:名称"`
	Scope  string `gorm:"size:200;comment:数据权限"`
	Weight string `gorm:"comment:权重"`
	base.Entity
}

func (*Menu) TableName() string {
	return "sys_menu"
}

// BeforeDelete 删除用户前清除用户与角色的关联信息
func (my *Menu) BeforeDelete(tx *gorm.DB) (err error) {
	if e := tx.Where("mid = ?", my.Id).Delete(&RoleMenu{}).Error; e != nil {
		err = fmt.Errorf("删除用户角色关联异常: <%s>", e.Error())
	}
	return
}
