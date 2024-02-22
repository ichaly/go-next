package sys

import "github.com/ichaly/go-next/pkg/base"

type RoleUser struct {
	base.Primary `mapstructure:",squash"`
	Rid          base.Id `gorm:"comment:角色ID"`
	Uid          base.Id `gorm:"comment:用户ID"`
}

func (RoleUser) TableName() string {
	return "sys_role_user"
}
