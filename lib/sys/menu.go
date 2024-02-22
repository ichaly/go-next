package sys

type Menu struct {
}

func (Menu) TableName() string {
	return "sys_menu"
}
