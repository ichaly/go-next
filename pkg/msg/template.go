package msg

// Template 模版
type Template struct {
}

func (*Template) TableName() string {
	return "msg_template"
}
