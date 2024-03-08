package cms

import (
	"github.com/ichaly/go-next/lib/base"
)

type Comment struct {
	Content           string `gorm:"type:text;comment:内容"`
	base.DeleteEntity `mapstructure:",squash"`
}

func (Comment) TableName() string {
	return "cms_comment"
}
