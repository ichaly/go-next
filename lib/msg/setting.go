package msg

// Setting 设置
type Setting struct {
}

func (*Setting) TableName() string {
	return "msg_setting"
}
