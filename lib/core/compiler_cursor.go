package core

import (
	"fmt"
	"github.com/vektah/gqlparser/v2/ast"
	"time"
)

func (my *compilerContext) appendCursorValue(field *ast.Field, child *ast.ChildValue) {
	sort := field.Arguments.ForName(SORT)
	//拼接原始条件
	if sort == nil {
		sort = &ast.Argument{Name: SORT, Value: &ast.Value{Kind: ast.ObjectValue, Children: []*ast.ChildValue{}}}
		field.Arguments = append(field.Arguments, sort)
	}
	sort.Value.Children = append(sort.Value.Children, child)
}

func (my *compilerContext) renderCursor(id int, field *ast.Field) {
	first := field.Arguments.ForName(FIRST)
	last := field.Arguments.ForName(LAST)
	if first == nil && last == nil {
		return
	}

	// 添加按主键排序的条件
	child := &ast.ChildValue{Name: "id", Value: &ast.Value{Kind: ast.EnumValue, Raw: "DESC"}}
	my.appendCursorValue(field, child)

	my.Write(`, CONCAT('`)
	my.Write(fmt.Sprintf("gj/%x:", time.Now().Unix()))
	my.Write(`', CONCAT_WS(',', `)
	my.Write(id)

	sort := field.Arguments.ForName(SORT)
	for i := 0; i < len(sort.Value.Children); i++ {
		my.Write(`, MAX(__cur_`)
		my.Write(i)
		my.Write(`)`)
	}

	my.Write(`)) as __cursor`)
}

func (my *compilerContext) renderCursorExclude(field *ast.Field) {
	sort := field.Arguments.ForName(SORT)
	if sort == nil {
		return
	}
	for i := 0; i < len(sort.Value.Children); i++ {
		my.Write(`- '__cur_`)
		my.Write(i)
		my.Write(`'`)
	}
}

func (my *compilerContext) renderCursorSelect(field *ast.Field) {
	sort := field.Arguments.ForName(SORT)
	if sort == nil {
		return
	}
	for i := 0; i < len(sort.Value.Children); i++ {
		my.Write(`, '__cur_`)
		my.Write(i)
		my.Write(`'`)
	}
}
