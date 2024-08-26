package core

import (
	"github.com/iancoleman/strcase"
	"github.com/ichaly/go-next/lib/util"
	"strings"
)

func (my *Metadata) Named(c Column) (string, string, string, string) {
	currentTable, currentColumn := c.Table, strings.TrimSuffix(c.Name, "_id")
	foreignTable, foreignColumn := c.TableRelation, c.ColumnRelation

	//移除前缀
	if val, ok := util.StartWithAny(currentTable, my.cfg.Prefixes...); ok {
		currentTable = strings.TrimPrefix(currentTable, val)
	}
	if val, ok := util.StartWithAny(foreignTable, my.cfg.Prefixes...); ok {
		foreignTable = strings.TrimPrefix(foreignTable, val)
	}

	//标准化列名
	if "id" == foreignColumn {
		foreignColumn = currentTable
	}
	foreignColumn = foreignColumn + "_list"

	//是否启用驼峰命名
	if my.cfg.UseCamel {
		currentTable = strcase.ToCamel(currentTable)
		currentColumn = strcase.ToLowerCamel(currentColumn)
		foreignTable = strcase.ToCamel(foreignTable)
		foreignColumn = strcase.ToLowerCamel(foreignColumn)
	}

	//是否是自关联
	if c.TableRelation == c.Table {
		currentColumn = "parent"
		foreignColumn = "children"
	}

	return currentTable, currentColumn, foreignTable, foreignColumn
}

func (my *Metadata) NamedTable(table string) string {
	//移除前缀
	if val, ok := util.StartWithAny(table, my.cfg.Prefixes...); ok {
		table = strings.TrimPrefix(table, val)
	}

	//是否启用驼峰命名
	if my.cfg.UseCamel {
		table = strcase.ToCamel(table)
	}

	return table
}

func (my *Metadata) NamedColumn(column string, ops ...ColumnOption) string {
	//应用配置选项
	for _, o := range ops {
		column = o(column)
	}

	//是否启用驼峰命名
	if my.cfg.UseCamel {
		column = strcase.ToLowerCamel(column)
	}

	return column
}

type ColumnOption func(s string) string

func WithTrimSuffix() ColumnOption {
	return func(s string) string {
		return strings.TrimSuffix(s, "_id")
	}
}

func JoinListSuffix() ColumnOption {
	return func(s string) string {
		return strings.Join([]string{s, "list"}, "_")
	}
}

func WithRecursion() ColumnOption {
	return func(s string) string {
		return strings.Join([]string{s, "list"}, "_")
	}
}
