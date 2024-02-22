package sys

import "github.com/ichaly/go-next/pkg/base"

type RoleMenu struct {
	base.Primary `mapstructure:",squash"`
	Rid          base.Id `gorm:"comment:角色ID"`
	Mid          base.Id `gorm:"comment:菜单ID"`
}

func (RoleMenu) TableName() string {
	return "sys_role_menu"
}
