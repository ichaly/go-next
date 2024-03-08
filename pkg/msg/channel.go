package msg

// Channel 渠道
type Channel struct {
}

func (*Channel) TableName() string {
	return "msg_channel"
}
