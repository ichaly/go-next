package core

import (
	"github.com/duke-git/lancet/v2/condition"
	"github.com/iancoleman/strcase"
	"github.com/ichaly/go-next/lib/core/internal"
	"github.com/ichaly/go-next/lib/util"
	"strings"
)

func (my *Metadata) Named(class, field string, ops ...NamedOption) (string, string) {
	//移除表前缀
	if val, ok := util.StartWithAny(class, my.cfg.Prefixes...); ok {
		class = strings.TrimPrefix(class, val)
	}

	//应用配置选项
	for _, o := range ops {
		field = o(class, field)
	}

	//是否驼峰命名
	if my.cfg.UseCamel {
		class = strcase.ToCamel(class)
		field = strcase.ToLowerCamel(field)
	}

	return class, field
}

type NamedOption func(table, column string) string

// WithTrimSuffix 移除`_id`后缀
func WithTrimSuffix() NamedOption {
	return func(t, s string) string {
		return strings.TrimSuffix(s, "_id")
	}
}

// JoinListSuffix 追加`_list`后缀
func JoinListSuffix() NamedOption {
	return func(t, s string) string {
		return strings.Join([]string{s, "list"}, "_")
	}
}

// SwapPrimaryKey 替换id列的名称
func SwapPrimaryKey(table string) NamedOption {
	return func(t, s string) string {
		if s == "id" {
			s = table
		}
		return s
	}
}

// NamedRecursion 重命名递归关联列名
func NamedRecursion(c *internal.Record, b bool) NamedOption {
	return func(t, s string) string {
		if c.TableRelation == c.TableName {
			s = condition.TernaryOperator(b, "parent", "children")
		}
		return s
	}
}
