package sys

import (
	"database/sql/driver"
	"github.com/ichaly/go-next/pkg/base"
)

type OauthKind string

const (
	Email  OauthKind = "email"
	Mobile OauthKind = "mobile"
	WeiXin OauthKind = "weixin"
)

func (my *OauthKind) Scan(value interface{}) error {
	*my = OauthKind(value.(string))
	return nil
}

func (my OauthKind) Value() (driver.Value, error) {
	return string(my), nil
}

type Oauth struct {
	Uid   base.ID   `gorm:"primaryKey;autoIncrement:false;comment:用户ID" json:",omitempty"`
	Kind  OauthKind `gorm:"primaryKey;size:50;comment:类型"`
	Value string    `gorm:"primaryKey;size:100;comment:值"`
	base.Entity
}

func (Oauth) TableName() string {
	return "sys_oauth"
}
