package msg

import "github.com/ichaly/go-next/pkg/base"

// Room 聊天室
type Room struct {
	Title        string `gorm:"size:200;comment:群组名称"`
	Avatar       string `gorm:"size:500;comment:群组头像"`
	Notice       string `gorm:"type:text;comment:群公告"`
	base.Entity  `mapstructure:",squash"`
	base.Deleted `mapstructure:",squash"`
}

func (*Room) TableName() string {
	return "msg_room"
}
