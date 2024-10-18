package core

import (
	"github.com/vektah/gqlparser/v2/ast"
	"strings"
)

func (my *compilerContext) appendWhereValue(field *ast.Field, value *ast.Value) {
	where := field.Arguments.ForName(WHERE)
	//拼接原始条件
	if where == nil {
		where = &ast.Argument{Name: WHERE, Value: value}
		field.Arguments = append(field.Arguments, where)
	} else {
		//使用AND包装拼接关联关系查询条件
		where.Value = &ast.Value{Kind: ast.ObjectValue, Children: []*ast.ChildValue{
			{Name: AND, Value: &ast.Value{Kind: ast.ListValue, Children: []*ast.ChildValue{{Value: where.Value}, {Value: value}}}},
		}}
	}
}

func (my *compilerContext) renderWhereField(field *ast.Field) {
	where := field.Arguments.ForName(WHERE)
	if where != nil {
		my.Write(` WHERE (`)
		my.renderWhereValue(where.Value)
		my.Write(`)`)
	}
}

func (my *compilerContext) renderWhereValue(value *ast.Value) {
	if value == nil {
		return
	}
	if value.Raw != "" {
		// TODO:使用?占位符,利用预编译提高性能
		if value.Kind == ast.EnumValue && value.Raw == "NOT_NULL" {
			my.Write("NOT NULL")
		} else if value.Kind == ast.StringValue {
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
		case IS, IN, EQ, NE, GT, GE, LT, LE, LIKE, I_LIKE, REGEX, I_REGEX:
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
			// TODO：更合适的办法？如果Definition为空，则认为是多表关联条件使用字段名称
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
