package core

import (
	"github.com/iancoleman/strcase"
	"github.com/ichaly/go-next/lib/util"
	"strings"
)

func (my *Metadata) Named(v Column) (string, string, string, string) {
	primaryTable, primaryColumn := v.Table, strings.TrimSuffix(v.Name, "_id")
	foreignTable, foreignColumn := v.TableRelation, v.ColumnRelation

	//移除前缀
	if val, ok := util.StartWithAny(primaryTable, my.cfg.Prefixes...); ok {
		primaryTable = strings.TrimPrefix(primaryTable, val)
	}
	if val, ok := util.StartWithAny(foreignTable, my.cfg.Prefixes...); ok {
		foreignTable = strings.TrimPrefix(foreignTable, val)
	}

	//标准化列名
	if "id" == foreignColumn {
		foreignColumn = primaryTable
	}
	foreignColumn = foreignColumn + "_list"

	//是否启用驼峰命名
	if my.cfg.UseCamel {
		primaryTable = strcase.ToCamel(primaryTable)
		primaryColumn = strcase.ToLowerCamel(primaryColumn)
		foreignTable = strcase.ToCamel(foreignTable)
		foreignColumn = strcase.ToLowerCamel(foreignColumn)
	}

	//是否是自关联
	if v.TableRelation == v.Table {
		primaryColumn = "parent"
		foreignColumn = "children"
	}

	return primaryTable, primaryColumn, foreignTable, foreignColumn
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

func (my *Metadata) NamedColumn(column string) string {
	column = column + "_list"

	//是否启用驼峰命名
	if my.cfg.UseCamel {
		column = strcase.ToLowerCamel(column)
	}

	return column
}
