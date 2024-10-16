package core

import (
	"github.com/vektah/gqlparser/v2/ast"
	"strings"
)

func (my *compilerContext) renderWhereValue(value *ast.Value) {
	if value == nil {
		return
	}
	if value.Raw != "" {
		// TODO:使用?占位符,利用预编译提高性能
		if value.Kind == ast.StringValue {
			my.Write("'")
			my.Write(value.Raw)
			my.Write("'")
		} else {
			my.Write(value.Raw)
		}
		return
	}
	for _, v := range value.Children {
		switch v.Name {
		case "is", "in", "eq", "ne", "gt", "ge", "lt", "le", "like", "iLike", "regex", "iRegex":
			if s, ok := dictionary[v.Name]; ok {
				my.Write(" ")
				my.Write(strings.ToUpper(s.Text))
				my.Write(" ")
			}
			my.renderWhereValue(v.Value)
		case OR, AND:
			my.Write("(")
			for i, child := range v.Value.Children {
				if i > 0 {
					my.Write(" ")
					my.Write(strings.ToUpper(v.Name))
					my.Write(" ")
				}
				my.renderWhereValue(child.Value)
			}
			my.Write(")")
		case NOT:
			my.Write("NOT (")
			my.renderWhereValue(value.Children[0].Value)
			my.Write(")")
		default:
			my.Write("(")
			// 如果Definition为空，则认为是多表关联条件使用字段名称
			// TODO：更合适的办法？
			if value.Definition != nil {
				name := strings.TrimSuffix(value.Definition.Name, SUFFIX_WHERE_INPUT)
				table, _ := my.meta.TableName(name, false)
				column, _ := my.meta.ColumnName(name, v.Name, false)
				my.Quoted(table)
				my.Write(".")
				my.Quoted(column)
			} else if v.Name != "" {
				my.Write("(")
				my.Write(v.Name)
				my.Write(")")
			}
			my.renderWhereValue(v.Value)
			my.Write(")")
		}
	}
}
