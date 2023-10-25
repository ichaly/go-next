package sys

import (
	"database/sql/driver"
	"github.com/ichaly/go-next/pkg/base"
)

type BindType string

const (
	Email  BindType = "email"
	Mobile BindType = "mobile"
	WeiXin BindType = "weixin"
)

func (my *BindType) Scan(value interface{}) error {
	*my = BindType(value.(string))
	return nil
}

func (my BindType) Value() (driver.Value, error) {
	return string(my), nil
}

type Bind struct {
	Uid   base.ID  `gorm:"primaryKey;autoIncrement:false;comment:用户ID" json:",omitempty"`
	Type  BindType `gorm:"primaryKey;size:50;comment:类型"`
	Value string   `gorm:"primaryKey;size:100;comment:值"`
	base.Entity
}

func (Bind) TableName() string {
	return "sys_bind"
}
