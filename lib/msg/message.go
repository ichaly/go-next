package msg

import "github.com/ichaly/go-next/pkg/base"

const (
	// Single 私聊
	Single MessageKind = iota
	// Multi 群聊
	Multi
	// Broadcast 广播
	Broadcast
	// System 系统
	System
	// Remind 提醒
	Remind
)

type MessageKind int8

type Message struct {
	Kind        MessageKind `gorm:"comment:消息类型"`
	Title       string      `gorm:"size:200;comment:标题"`
	Cover       string      `gorm:"type:text;comment:封面"`
	Content     string      `gorm:"type:text;comment:内容"`
	FromId      *base.Id    `gorm:"comment:发送人"` //系统消息的发送者也是一个用户,快递消息、支付通知、活动精选、系统消息、会员动态
	UserId      base.Id     `gorm:"comment:接收人"` //私聊是用户ID,群聊是群ID
	SourceId    base.Id     `gorm:"comment:关联ID"`
	SourceType  int8        `gorm:"comment:关联类型"` //项目、文章、评论、商品、视频、图片、用户
	base.Entity `mapstructure:",squash"`
}

func (*Message) TableName() string {
	return "msg_message"
}

// 参考http://117.50.186.19:3000/index.html#/crud/new
