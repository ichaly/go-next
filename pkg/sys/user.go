package sys

import (
	"fmt"
	"github.com/ichaly/go-next/lib/base"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Username    string     `gorm:"size:200;comment:名称;index:,unique" validate:"required"`
	Password    string     `gorm:"size:200;comment:密码;" gql:"-"`
	Nickname    string     `gorm:"size:50;comment:昵称"` //https://github.com/DanPlayer/randomname/tree/main
	Birthday    *time.Time `gorm:"comment:生日"`
	Avatar      string     `gorm:"size:200;comment:头像"` //https://golangnote.com/topic/274.html
	Source      string     `gorm:"comment:来源"`
	base.Entity `mapstructure:",squash"`
}

func (*User) TableName() string {
	return "sys_user"
}

func (*User) Description() string {
	return "用户信息"
}

// BeforeDelete 删除用户前清除用户与角色的关联信息
func (my *User) BeforeDelete(tx *gorm.DB) (err error) {
	// 清除casbin用户与角色关联
	if _, e := roleService.DeleteUser(my.Id); e != nil {
		err = fmt.Errorf("清除casbin角色权限异常: <%s>", e.Error())
	}
	if e := tx.Where("uid = ?", my.Id).Delete(&RoleUser{}).Error; e != nil {
		err = fmt.Errorf("删除用户角色关联异常: <%s>", e.Error())
	}
	return
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
