package auth

import (
	"context"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/ichaly/go-next/pkg/base"
	"github.com/ichaly/go-next/pkg/util"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ClientStore struct {
	db     *gorm.DB
	config *base.Config
}

func NewOauthClientStore(d *gorm.DB, c *base.Config) oauth2.ClientStore {
	client := &Client{}
	if !d.Migrator().HasTable(client.TableName()) {
		if err := d.Migrator().CreateTable(client); err != nil {
			panic(err)
		}
	}
	return &ClientStore{db: d, config: c}
}

func (my *ClientStore) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {
	c := Client{Passkey: my.config.Oauth.Passkey}
	e := my.db.WithContext(ctx).Model(c).Where("id = ?", id).Take(&c).Error
	return &c, e
}

type Client struct {
	Passkey string `gorm:"-"`
	Domain  string `gorm:"type:varchar(512)"`
	Secret  string `gorm:"type:varchar(512)"`
	Public  bool
	base.Entity
}

func (Client) TableName() string {
	return "sys_client"
}

func (my *Client) BeforeCreate(tx *gorm.DB) error {
	return my.encryptSecret(tx)
}

func (my *Client) BeforeUpdate(tx *gorm.DB) error {
	return my.encryptSecret(tx)
}

func (my *Client) encryptSecret(tx *gorm.DB) error {
	if my.Secret == "" {
		return nil
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(my.Secret), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	tx.Statement.SetColumn("Secret", string(hash))
	return nil
}

func (my *Client) GetID() string {
	return util.FormatLong(int64(my.ID))
}

func (my *Client) GetSecret() string {
	return my.Secret
}

func (my *Client) GetDomain() string {
	return my.Domain
}

func (my *Client) IsPublic() bool {
	return my.Public
}

func (my *Client) GetUserID() string {
	return ""
}

func (my *Client) VerifyPassword(secret string) bool {
	if my.Passkey == secret {
		return true
	}
	err := bcrypt.CompareHashAndPassword([]byte(my.Secret), []byte(secret))
	return err == nil
}
