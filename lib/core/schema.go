package core

import (
	_ "embed"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/ichaly/go-next/lib/util"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

//go:embed sql/pgsql.sql
var pgsql string

type Table struct {
	Name        string
	Prefix      string
	Description string
	Columns     []*Column
}

type Column struct {
	Name        string `mapstructure:""`
	Type        string
	Description string
	IsPrimary   bool
	IsForeign   bool
	IsNullable  bool
}

type Schema struct {
	Nodes map[string]*Table
}

func NewSchema(db *gorm.DB) (*string, error) {
	var list []*struct {
		Name        string `gorm:"column:column_name;"`
		Type        string `gorm:"column:data_type;"`
		Table       string `gorm:"column:table_name;"`
		Comment     string `gorm:"column:table_description;"`
		IsPrimary   bool   `gorm:"column:is_primary;"`
		IsForeign   bool   `gorm:"column:is_foreign;"`
		IsNullable  bool   `gorm:"column:is_nullable;"`
		Description string `gorm:"column:column_description;"`
	}
	err := db.Raw(pgsql).Scan(&list).Error
	if err != nil {
		return nil, err
	}
	schema := &Schema{
		Nodes: make(map[string]*Table),
	}
	for _, v := range list {
		var c Column
		err := mapstructure.Decode(v, &c)
		if err != nil {
			continue
		}
		node, ok := schema.Nodes[v.Table]
		if !ok {
			name := strcase.ToCamel(v.Table)
			node = &Table{
				Name:        name,
				Description: v.Comment,
				Columns:     make([]*Column, 0),
			}
			schema.Nodes[name] = node
		}
		node.Columns = append(node.Columns, &c)
	}
	json, err := util.MarshalJson(schema)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%v\n", json)
	return nil, err
}
