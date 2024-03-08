package cms

import (
	"github.com/ichaly/go-next/lib/base"
)

type Content struct {
	Title             string `gorm:"size:200;comment:标题"`
	Intro             string `gorm:"size:1000;comment:简介"`
	Content           string `gorm:"type:text;comment:内容"`
	Source            string `gorm:"comment:来源"`
	Kind              Kind   `gorm:"comment:类型"`
	base.DeleteEntity `mapstructure:",squash"`
}

func (Content) TableName() string {
	return "cms_content"
}

func (Content) Description() string {
	return "内容信息"
}
