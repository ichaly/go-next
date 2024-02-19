package msg

import "github.com/ichaly/go-next/pkg/base"

// 参考http://117.50.186.19:3000/index.html#/crud/new
// Session 会话按照MessageType分类
type Session struct {
	Title       string      `gorm:"size:200;comment:会话标题"`
	Avatar      string      `gorm:"size:500;comment:会话头像"`
	Content     string      `gorm:"type:text;comment:最后内容"`
	FromId      base.Id     `gorm:"comment:来源ID"`
	ToId        base.Id     `gorm:"comment:接收人"`
	Unread      int8        `gorm:"comment:未读数"`
	MessageType MessageKind `gorm:"comment:消息类型"`
	base.Entity `mapstructure:",squash"`
}

func (*Session) TableName() string {
	return "msg_session"
}
