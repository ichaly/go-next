package msg

import "github.com/ichaly/go-next/pkg/base"

// Notice 通知
type Notice struct {
	Title             string `gorm:"size:200;comment:标题"`
	Content           string `gorm:"type:text;comment:内容"`
	base.DeleteEntity `mapstructure:",squash"`
}

func (*Notice) TableName() string {
	return "msg_notice"
}
