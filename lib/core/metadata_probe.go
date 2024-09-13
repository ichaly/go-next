package core

import (
	"github.com/duke-git/lancet/v2/condition"
	"github.com/duke-git/lancet/v2/maputil"
	"github.com/ichaly/go-next/lib/core/internal"
	"github.com/ichaly/go-next/lib/util"
)

type Input struct {
	Name string
	Type string
}

type Class struct {
	Name        string
	Description string
	Fields      map[string]Field
}

type Field struct {
	Type             string `gorm:"column:data_type;"`
	Name             string `gorm:"column:column_name;"`
	Table            string `gorm:"column:table_name;"`
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

var tableOption = func(my *Metadata) error {
	// 查询表结构
	if err := my.db.Raw(pgsql).Scan(&my.list).Error; err != nil {
		return err
	}

	my.tree, my.edge = make(internal.AnyMap[Class]), make(internal.AnyMap[internal.AnyMap[Field]])

	for _, c := range my.list {
		//判断是否包含黑名单关键字,执行忽略跳过
		if _, ok := util.ContainsAny(c.Name, my.cfg.BlockList...); ok {
			continue
		}
		if _, ok := util.ContainsAny(c.Table, my.cfg.BlockList...); ok {
			continue
		}

		//转化类型
		c.Type = condition.TernaryOperator(c.IsPrimary, "ID", my.cfg.Mapping[c.Type])

		//规范命名
		table, column := my.Named(c.Table, c.Name)

		//索引节点
		maputil.GetOrSet(my.tree, table, Class{
			Name:        c.Table,
			Description: c.TableDescription,
			Fields:      make(internal.AnyMap[Field]),
		}).Fields[column] = c

		//索引外键
		if c.IsForeign {
			maputil.GetOrSet(my.edge, table, make(map[string]Field))[column] = c
		}
	}

	return nil
}
