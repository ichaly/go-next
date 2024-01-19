package sys

import (
	"github.com/ichaly/go-next/pkg/base"
)

type Team struct {
	Name string `gorm:"size:200;comment:名称"`
	base.Entity
}

func (Team) TableName() string {
	return "sys_team"
}
