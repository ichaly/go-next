package sys

import (
	"github.com/ichaly/go-next/pkg/base"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Username    string     `gorm:"size:200;comment:名称;index:,unique" validate:"required"`
	Password    string     `gorm:"size:200;comment:密码;" gql:"-"`
	Nickname    string     `gorm:"size:50;comment:昵称"`
	Avatar      string     `gorm:"size:200;comment:头像"`
	Source      string     `gorm:"comment:来源"`
	Birthday    *time.Time `gorm:"comment:生日"`
	base.Entity `mapstructure:",squash"`
}

func (*User) TableName() string {
	return "sys_user"
}

func (*User) Description() string {
	return "用户信息"
}

func (my *User) BeforeCreate(tx *gorm.DB) error {
	return my.encryptPassword(tx)
}

func (my *User) BeforeUpdate(tx *gorm.DB) error {
	return my.encryptPassword(tx)
}

func (my *User) encryptPassword(tx *gorm.DB) error {
	if my.Password == "" {
		return nil
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(my.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	tx.Statement.SetColumn("Password", string(hash))
	return nil
}
