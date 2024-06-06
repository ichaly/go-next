package core

import (
	"fmt"
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
	_, _ = NewSchema(db)
}
