package sys

import (
	"database/sql"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (my *UserService) FindByUsername(username string) (User, error) {
	usr := User{}
	err := my.db.Model(&usr).Joins("left join sys_oauth on sys_oauth.uid = sys_user.id").
		Where("sys_user.username = @username or sys_oauth.value = @username", sql.Named("username", username)).
		First(&usr).Error
	return usr, err
}

func (my *UserService) Bind(u *User, k OauthKind) {
	my.db.Save(u)
	my.db.Save(&Oauth{Kind: k, Uid: u.ID, Value: u.Username})
}
