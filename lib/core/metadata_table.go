package core

import (
	"github.com/duke-git/lancet/v2/condition"
	"github.com/duke-git/lancet/v2/maputil"
	"github.com/ichaly/go-next/lib/core/internal"
	"github.com/ichaly/go-next/lib/util"
)

type Table struct {
	Name        string
	Description string
	Columns     Directory[Column]
}

type Column struct {
	Type             string `gorm:"column:data_type;"`
	Name             string `gorm:"column:column_name;"`
	Table            string `gorm:"column:table_name;"`
	IsPrimary        bool   `gorm:"column:is_primary;"`
	IsForeign        bool   `gorm:"column:is_foreign;"`
	IsNullable       bool   `gorm:"column:is_nullable;"`
	Description      string `gorm:"column:column_description;"`
	TableRelation    string `gorm:"column:table_relation;"`
	ColumnRelation   string `gorm:"column:column_relation;"`
	TableDescription string `gorm:"column:table_description;"`
}

func (my Column) SetType(dataType string) Column {
	my.Type = dataType
	return my
}

func (my *Metadata) load() error {
	// 查询表结构
	tx := my.db.Raw(pgsql).Scan(&my.list)
	if err := tx.Error; err != nil {
		return err
	}

	my.tree = make(Directory[Table])
	my.keys = make(Directory[Directory[Column]])

	for _, c := range my.list {
		//判断是否包含黑名单关键字,执行忽略跳过
		if _, ok := util.ContainsAny(c.Name, my.cfg.BlockList...); ok {
			continue
		}
		if _, ok := util.ContainsAny(c.Table, my.cfg.BlockList...); ok {
			continue
		}

		//转化类型
		c.Type = condition.TernaryOperator(c.IsPrimary, "ID", internal.DataTypes[c.Type])

		//规范命名
		table, column := my.Named(c.Table, c.Name)

		//索引节点
		maputil.GetOrSet(my.tree, table, Table{
			Name:        c.Table,
			Description: c.TableDescription,
			Columns:     make(Directory[Column]),
		}).Columns[column] = c

		//索引外键
		if c.IsForeign {
			maputil.GetOrSet(my.keys, table, make(map[string]Column))[column] = c
		}
	}

	return nil
}
