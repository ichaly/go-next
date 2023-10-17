package cms

import (
	"github.com/ichaly/go-next/pkg/base"
)

type Model struct {
	base.Entity `mapstructure:",squash"`
}

func (Model) TableName() string {
	return "cms_model"
}
