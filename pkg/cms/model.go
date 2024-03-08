package cms

import (
	"github.com/ichaly/go-next/lib/base"
)

type Model struct {
	base.Entity `mapstructure:",squash"`
}

func (Model) TableName() string {
	return "cms_model"
}
