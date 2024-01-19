package sys

import "github.com/ichaly/go-next/pkg/base"

type Role struct {
	Code   string `gorm:"index:,unique;size:200;comment:编码"`
	Name   string `gorm:"size:200;comment:名称"`
	Scope  string `gorm:"size:200;comment:数据权限"`
	Weight string `gorm:"comment:权重"`
	base.Entity
}

func (Role) TableName() string {
	return "sys_role"
}
