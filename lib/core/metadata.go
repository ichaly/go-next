package core

import (
	_ "embed"
	"github.com/iancoleman/strcase"
	"github.com/ichaly/go-next/lib/core/internal"
	"github.com/ichaly/go-next/lib/util"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"strings"
)

//go:embed assets/pgsql.sql
var pgsql string

//go:embed assets/schema.tpl
var schema string

type Table struct {
	Name        string
	Original    string
	Description string
	Columns     []*Column
}

type Column struct {
	Name        string
	Type        string
	Description string
	IsPrimary   bool
	IsForeign   bool
	IsNullable  bool
}

type Metadata struct {
	db    *gorm.DB
	cfg   *internal.TableConfig
	Nodes map[string]*Table
}

func NewMetadata(d *gorm.DB, v *viper.Viper) (*Metadata, error) {
	cfg := &internal.TableConfig{Mapping: internal.DataTypes}
	if err := v.Sub("schema").Unmarshal(cfg); err != nil {
		return nil, err
	}
	metadata := &Metadata{
		db: d, cfg: cfg, Nodes: make(map[string]*Table),
	}
	if err := metadata.load(); err != nil {
		return nil, err
	}
	return metadata, nil
}

func (my *Metadata) load() error {
	var list []*struct {
		Name             string `gorm:"column:column_name;"`
		Type             string `gorm:"column:data_type;"`
		Table            string `gorm:"column:table_name;"`
		IsPrimary        bool   `gorm:"column:is_primary;"`
		IsForeign        bool   `gorm:"column:is_foreign;"`
		IsNullable       bool   `gorm:"column:is_nullable;"`
		Description      string `gorm:"column:column_description;"`
		TableDescription string `gorm:"column:table_description;"`
	}
	if err := my.db.Raw(pgsql).Scan(&list).Error; err != nil {
		return err
	}
	for _, v := range list {
		//判断是否包含黑名单关键字,执行忽略跳过
		if _, ok := util.ContainsAny(v.Name, my.cfg.BlockList...); ok {
			continue
		}
		if _, ok := util.ContainsAny(v.Table, my.cfg.BlockList...); ok {
			continue
		}

		//解析列
		var c Column
		if err := mapstructure.Decode(v, &c); err != nil {
			return err
		}

		//解析表
		name := v.Table
		if val, yes := util.StartWithAny(name, my.cfg.Prefixes...); yes {
			name = strings.Replace(name, val, "", 1)
		}
		if my.cfg.UseCamel {
			name = strcase.ToCamel(name)
			c.Name = strcase.ToLowerCamel(c.Name)
		}

		//索引节点
		if node, ok := my.Nodes[name]; !ok {
			node = &Table{
				Name:        name,
				Original:    v.Table,
				Description: v.TableDescription,
				Columns:     make([]*Column, 0),
			}
			my.Nodes[name] = node
		} else {
			node.Columns = append(node.Columns, &c)
		}
	}
	return nil
}