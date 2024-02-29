package sys

import (
	"fmt"
	"github.com/ichaly/go-next/pkg/base"
	"gorm.io/gorm"
)

type RoleUser struct {
	base.Primary `mapstructure:",squash"`
	RoleId       base.Id `gorm:"comment:角色ID"`
	UserId       base.Id `gorm:"comment:用户ID"`
}

func (*RoleUser) TableName() string {
	return "sys_role_user"
}

// AfterCreate 添加关联后，创建casbin的用户与角色的关联
func (my *RoleUser) AfterCreate(tx *gorm.DB) (err error) {
	if _, e := roleService.AddUserPolicy(my.UserId, my.RoleId); e != nil {
		err = fmt.Errorf("关联用户和角色到casbin异常: <%s>", e.Error())
	}
	return
}
