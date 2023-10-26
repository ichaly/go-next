package sys

import (
	"database/sql/driver"
	"github.com/ichaly/go-next/pkg/base"
)

type OauthType string

const (
	Email  OauthType = "email"
	Mobile OauthType = "mobile"
	WeiXin OauthType = "weixin"
)

func (my *OauthType) Scan(value interface{}) error {
	*my = OauthType(value.(string))
	return nil
}

func (my OauthType) Value() (driver.Value, error) {
	return string(my), nil
}

type Oauth struct {
	Uid   base.ID   `gorm:"primaryKey;autoIncrement:false;comment:用户ID" json:",omitempty"`
	Type  OauthType `gorm:"primaryKey;size:50;comment:类型"`
	Value string    `gorm:"primaryKey;size:100;comment:值"`
	base.Entity
}

func (Oauth) TableName() string {
	return "sys_oauth"
}
