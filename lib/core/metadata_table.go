package core

import (
	"github.com/duke-git/lancet/v2/condition"
	"github.com/duke-git/lancet/v2/maputil"
	"github.com/ichaly/go-next/lib/core/internal"
	"github.com/ichaly/go-next/lib/util"
	"github.com/mitchellh/mapstructure"
)

type Input struct {
	Name string
	Type string
}

type Class struct {
	Name        string
	Description string
	Fields      Map[Field]
}

type Field struct {
	Type             string `gorm:"column:data_type;"`
	Name             string `gorm:"column:column_name;"`
	TableName        string `gorm:"column:table_name;"`
	IsPrimary        bool   `gorm:"column:is_primary;"`
	IsForeign        bool   `gorm:"column:is_foreign;"`
	IsNullable       bool   `gorm:"column:is_nullable;"`
	IsIterable       bool   `gorm:"column:is_iterable;"`
	Description      string `gorm:"column:column_description;"`
	TableRelation    string `gorm:"column:table_relation;"`
	ColumnRelation   string `gorm:"column:column_relation;"`
	TableDescription string `gorm:"column:table_description;"`
}

func (my Field) SetType(dataType string) Field {
	my.Type = dataType
	return my
}

func (my *Metadata) load() error {
	// 查询表结构
	if err := my.db.Raw(pgsql).Scan(&my.list).Error; err != nil {
		return err
	}

	my.tree, my.keys = make(Map[Class]), make(Map[Map[Field]])

	for _, c := range my.list {
		//判断是否包含黑名单关键字,执行忽略跳过
		if _, ok := util.ContainsAny(c.Name, my.cfg.BlockList...); ok {
			continue
		}
		if _, ok := util.ContainsAny(c.TableName, my.cfg.BlockList...); ok {
			continue
		}

		//转化类型
		c.Type = condition.TernaryOperator(c.IsPrimary, "ID", internal.DataTypes[c.Type])

		//规范命名
		table, column := my.Named(c.TableName, c.Name)

		//索引节点
		maputil.GetOrSet(my.tree, table, Class{
			Name:        c.TableName,
			Description: c.TableDescription,
			Fields:      make(Map[Field]),
		}).Fields[column] = c

		//索引外键
		if c.IsForeign {
			maputil.GetOrSet(my.keys, table, make(map[string]Field))[column] = c
		}
	}

	if err := mapstructure.Decode(my.tree, &my.Nodes); err != nil {
		return err
	}

	return nil
}
