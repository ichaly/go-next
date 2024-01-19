package main

import (
	"fmt"
	"github.com/ichaly/go-next/lib/sys"
	"github.com/ichaly/go-next/pkg/auth"
	"github.com/ichaly/go-next/pkg/base"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"sync"
	"testing"
	"time"
)

type Demo struct {
	Name        string `gorm:"size:200;comment:名字"`
	base.Entity `mapstructure:",squash"`
}

func (*Demo) TableName() string {
	return "sys_demo"
}

func TestConnect(t *testing.T) {
	cfg, err := base.NewConfig("../cfg/dev.yml")
	if err != nil {
		t.Error(err)
	}
	d := &Demo{}
	db, err := base.NewConnect(cfg, []gorm.Plugin{base.NewSonyFlake()}, []interface{}{d})
	if err != nil {
		t.Error(err)
	}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(n int) {
			err := db.Transaction(func(tx *gorm.DB) error {
				var demos []*Demo
				_tx := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Model(&Demo{}).Where("state = ?", 0)
				_tx.Find(&demos)
				var ids []base.ID
				for _, d := range demos {
					ids = append(ids, d.ID)
				}
				time.Sleep(50 * time.Millisecond)
				return tx.Model(&Demo{}).Save(&Demo{Name: fmt.Sprintf("用户%d", n)}).Error
			})
			if err != nil {
				t.Error(err)
			}

			sqDb, _ := db.DB()
			t.Log("in use :", sqDb.Stats().InUse)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func TestUser(t *testing.T) {
	cfg, err := base.NewConfig("../cfg/dev.yml")
	if err != nil {
		t.Error(err)
	}
	db, err := base.NewConnect(cfg, []gorm.Plugin{base.NewSonyFlake()}, []interface{}{})
	if err != nil {
		t.Error(err)
	}
	err = db.Model(&sys.User{}).Save(&sys.User{Username: "admin", Password: "123456", Nickname: "管理员"}).Error
	if err != nil {
		t.Error(err)
	}
}

func TestClient(t *testing.T) {
	cfg, err := base.NewConfig("../cfg/dev.yml")
	if err != nil {
		t.Error(err)
	}
	db, err := base.NewConnect(cfg, []gorm.Plugin{base.NewSonyFlake()}, []interface{}{})
	if err != nil {
		t.Error(err)
	}
	//secret := strings.ReplaceAll(uuid.Must(uuid.NewRandom()).String(), "-", "")
	//t.Log(secret)
	client := auth.Client{Secret: "90c5f47a4b1e42f48efb20e0ed30cae7", Domain: "http://127.0.0.1:8080/oauth/callback/go-next"}
	client.ID = 12708495015018519
	err = db.Save(&client).Error
	if err != nil {
		t.Error(err)
	}
}
