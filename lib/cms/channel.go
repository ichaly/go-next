package cms

import (
	"github.com/ichaly/go-next/pkg/base"
)

type Channel struct {
	Name        string `gorm:"size:200;comment:名称"`
	base.Entity `mapstructure:",squash"`
}

func (Channel) TableName() string {
	return "cms_channel"
}
