package core

import (
	_ "embed"
	"github.com/iancoleman/strcase"
	"github.com/ichaly/go-next/lib/core/internal"
	"github.com/ichaly/go-next/lib/util"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"regexp"
	"strings"
)

//go:embed assets/pgsql.sql
var pgsql string

//go:embed assets/schema.tpl
var schema string

var dbTypeRe = regexp.MustCompile(`([a-zA-Z ]+)(\((.+)\))?`)

type Table struct {
	Name        string
	Prefix      string
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
	cfg := &internal.TableConfig{}
	if err := v.Unmarshal(cfg); err != nil {
		return nil, err
	}
	metadata := &Metadata{
		d, cfg, make(map[string]*Table),
	}
	if err := metadata.load(); err != nil {
		return nil, err
	}
	return metadata, nil
}

func (my *Metadata) load() error {
	var list []*struct {
		Name              string `gorm:"column:column_name;"`
		Type              string `gorm:"column:data_type;"`
		Table             string `gorm:"column:table_name;"`
		IsPrimary         bool   `gorm:"column:is_primary;"`
		IsForeign         bool   `gorm:"column:is_foreign;"`
		IsNullable        bool   `gorm:"column:is_nullable;"`
		TableDescription  string `gorm:"column:table_description;"`
		ColumnDescription string `gorm:"column:column_description;"`
	}
	if err := my.db.Raw(pgsql).Scan(&list).Error; err != nil {
		return err
	}
	metadata := &Metadata{
		Nodes: make(map[string]*Table),
	}
	for _, v := range list {
		var c Column
		if err := mapstructure.Decode(v, &c); err != nil {
			return err
		}
		c.Name = strcase.ToCamel(c.Name)
		node, ok := metadata.Nodes[v.Table]
		if !ok {
			node = &Table{
				Name:        v.Table,
				Columns:     make([]*Column, 0),
				Description: v.TableDescription,
			}
			if val, yes := util.StartWithAny(v.Table, my.cfg.Prefixes...); yes {
				node.Name = strings.Replace(v.Table, val, "", 1)
				node.Prefix = val
			}
			node.Name = strcase.ToCamel(node.Name)
			metadata.Nodes[node.Name] = node
		}
		node.Columns = append(node.Columns, &c)
	}
	return nil
}
