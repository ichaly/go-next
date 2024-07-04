package core

import (
	"fmt"
	"github.com/ichaly/go-next/lib/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

func TestDataSource(t *testing.T) {
	args := []interface{}{"postgres", "postgres", "127.0.0.1", 5678, "gcms"}
	db, err := gorm.Open(postgres.Open(fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s TimeZone=Asia/Shanghai", args...,
	)), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		t.Fatal(err)
	}
	metadata, err := NewMetadata(db)
	if err != nil {
		t.Fatal(err)
	}
	json, err := util.MarshalJson(metadata)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(json)
}

type Parent interface {
	GetName() string
}

type Child interface {
	GetAge() int
}

type Son struct {
}

func (s *Son) GetName() string {
	return "son"
}

//func (s *Son) GetAge() int {
//	return 18
//}

func NewParent(p Parent) *Child {
	return nil
}

func TestEmbedded(t *testing.T) {
	var son *Son
	NewParent(son)
}
