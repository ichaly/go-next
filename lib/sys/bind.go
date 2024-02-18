package sys

import (
	"database/sql/driver"
	"github.com/ichaly/go-next/pkg/base"
)

type BindKind string

func (my *BindKind) Scan(value interface{}) error {
	*my = BindKind(value.(string))
	return nil
}

func (my BindKind) Value() (driver.Value, error) {
	return string(my), nil
}

type Bind struct {
	Uid   base.Id  `gorm:"primaryKey;autoIncrement:false;comment:用户ID" json:",omitempty"`
	Kind  BindKind `gorm:"primaryKey;size:50;comment:类型"`
	Value string   `gorm:"primaryKey;size:100;comment:值"`
	base.Entity
}

func (Bind) TableName() string {
	return "sys_bind"
}
