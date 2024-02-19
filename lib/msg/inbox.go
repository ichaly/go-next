package msg

import (
	"github.com/ichaly/go-next/pkg/base"
	"time"
)

const (
	// Remind 提醒
	Remind MessageKind = iota
	// Message 消息
	Message
	// Notification 通知
	Notification
)

type MessageKind int8

// Inbox 消息
type Inbox struct {
	Title       string      `gorm:"size:200;comment:标题"`
	Content     string      `gorm:"type:text;comment:内容"`
	FromId      base.Id     `gorm:"comment:发送人"` //系统消息的发送者也是一个用户,快递消息、支付通知、活动精选、系统消息、会员动态
	ToId        base.Id     `gorm:"comment:接收人"`
	ReadAt      *time.Time  `gorm:"comment:已读时间" json:",omitempty"`
	SourceId    *base.Id    `gorm:"comment:关联ID"`
	SourceType  int8        `gorm:"comment:关联类型"` //项目、文章、评论、商品、视频、图片、用户
	MessageType MessageKind `gorm:"comment:消息类型"` //公告、私信、提醒、订阅、点赞、评论、关注、收藏
	base.Entity `mapstructure:",squash"`
}

func (*Inbox) TableName() string {
	return "msg_inbox"
}
