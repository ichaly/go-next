package core

import (
	"github.com/duke-git/lancet/v2/condition"
	"github.com/iancoleman/strcase"
	"github.com/ichaly/go-next/lib/util"
	"strings"
)

func (my *Metadata) Named(table, column string, ops ...ColumnOption) (string, string) {
	//移除表前缀
	if val, ok := util.StartWithAny(table, my.cfg.Prefixes...); ok {
		table = strings.TrimPrefix(table, val)
	}

	//应用配置选项
	for _, o := range ops {
		column = o(table, column)
	}

	//是否驼峰命名
	if my.cfg.UseCamel {
		table = strcase.ToCamel(table)
		column = strcase.ToLowerCamel(column)
	}

	return table, column
}

type ColumnOption func(table, column string) string

func WithTrimSuffix() ColumnOption {
	return func(t, s string) string {
		return strings.TrimSuffix(s, "_id")
	}
}

func PrimaryColumn(table string) ColumnOption {
	return func(t, s string) string {
		if s == "id" {
			s = table
		}
		return s
	}
}

func JoinListSuffix() ColumnOption {
	return func(t, s string) string {
		return strings.Join([]string{s, "list"}, "_")
	}
}

func WithRecursion(c Column, b bool) ColumnOption {
	return func(t, s string) string {
		if c.TableRelation == c.Table {
			s = condition.TernaryOperator(b, "parent", "children")
		}
		return s
	}
}
