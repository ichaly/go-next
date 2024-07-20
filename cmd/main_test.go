package main

import (
	"github.com/ichaly/go-next/lib/base"
)

type Demo struct {
	Name        string `gorm:"size:200;comment:名字"`
	base.Entity `mapstructure:",squash"`
}

func (*Demo) TableName() string {
	return "sys_demo"
}
