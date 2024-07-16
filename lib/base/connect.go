package base

import (
	"fmt"
	"github.com/ichaly/go-next/lib/base/internal"
	"github.com/ichaly/go-next/lib/gql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func NewConnect(v *viper.Viper, c *Config, p []gorm.Plugin, e []interface{}) (*gorm.DB, error) {
	cfg := &internal.DatabaseConfig{}
	if err := v.Sub("database").Unmarshal(c); err != nil {
		return nil, err
	}

	db, err := gorm.Open(
		buildDialect(&cfg.DataSource),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
	)
	if err != nil {
		return nil, err
	}
	for _, v := range p {
		err = db.Use(v)
		if err != nil {
			return nil, err
		}
	}
	if c.Debug {
		for _, v := range e {
			name, desc := "", ""
			if n, ok := v.(schema.Tabler); ok {
				name = n.TableName()
			}
			if n, ok := v.(gql.GqlDescription); ok {
				desc = n.Description()
			}
			options := fmt.Sprintf(";comment on table %s is '%s';", name, desc)
			tx := db
			if name != "" && desc != "" {
				tx = db.Set("gorm:table_options", options)
			}
			err = tx.AutoMigrate(v)
			if err != nil {
				return nil, err
			}
		}
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(5)
	sqlDb.SetMaxOpenConns(90)
	sqlDb.SetConnMaxLifetime(5 * time.Minute)
	return db, nil
}

func buildDialect(ds *internal.DataSource) gorm.Dialector {
	args := []interface{}{ds.Username, ds.Password, ds.Host, ds.Port, ds.Name}
	if ds.Dialect == "mysql" {
		return mysql.Open(fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", args...,
		))
	} else {
		return postgres.Open(fmt.Sprintf(
			"user=%s password=%s host=%s port=%d dbname=%s TimeZone=Asia/Shanghai", args...,
		))
	}
}