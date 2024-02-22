package msg

import (
	"github.com/ichaly/go-next/pkg/base"
)

// Inbox 消息
type Inbox struct {
	Mid          base.Id `gorm:"comment:消息ID"`
	Uid          base.Id `gorm:"comment:用户ID"`
	base.Entity  `mapstructure:",squash"`
	base.Deleted `mapstructure:",squash"`
}

func (*Inbox) TableName() string {
	return "msg_inbox"
}
