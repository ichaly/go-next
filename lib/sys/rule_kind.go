package sys

import (
	"database/sql/driver"
)

type RuleKind string

const (
	Path   RuleKind = "path"
	Menu   RuleKind = "menu"
	Action RuleKind = "action"
)

func (my *RuleKind) Scan(value interface{}) error {
	*my = RuleKind(value.(string))
	return nil
}

func (my *RuleKind) Value() (driver.Value, error) {
	return my, nil
}

func (*RuleKind) Description() string {
	return "权限类型"
}

func (*RuleKind) EnumValues() map[string]*struct {
	Value             interface{}
	Description       string
	DeprecationReason string
} {
	return map[string]*struct {
		Value             interface{}
		Description       string
		DeprecationReason string
	}{
		"PATH":   {Value: Path, Description: "目录"},
		"MENU":   {Value: Menu, Description: "菜单"},
		"ACTION": {Value: Action, Description: "动作"},
	}
}
