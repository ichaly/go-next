package msg

import (
	"github.com/ichaly/go-next/pkg/base"
	"time"
)

type Message struct {
	Title       string `gorm:"size:200;comment:标题"`
	Content     string `gorm:"type:text;comment:内容"`
	EventType   int8   `gorm:"comment:事件类型"`
	EntityType  int8   `gorm:"comment:实体类型"` //快递消息、支付通知、活动精选、系统消息、会员动态
	NoticeType  int8   `gorm:"comment:通知类型"` //系统消息、订阅消息
	base.Entity `mapstructure:",squash"`
}

type Session struct {
	Content     string  `gorm:"type:text;comment:通知内容"`
	MsgId       base.Id `gorm:"comment:消息ID"`
	FromId      base.Id `gorm:"comment:发送人"`
	ToId        base.Id `gorm:"comment:接收人"`
	base.Entity `mapstructure:",squash"`
}

type Receiver struct {
	MsgId       base.Id    `gorm:"comment:消息ID;" json:",omitempty"`
	UserId      base.Id    `gorm:"comment:用户ID;" json:",omitempty"`
	ReadAt      *time.Time `gorm:"comment:已读时间;" json:",omitempty"`
	base.Entity `mapstructure:",squash"`
}
