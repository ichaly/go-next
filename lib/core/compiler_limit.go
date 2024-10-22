package core

import (
	"github.com/vektah/gqlparser/v2/ast"
	"strconv"
)

func (my *compilerContext) renderLimit(field *ast.Field) {
	my.Write(` LIMIT `)
	var value *ast.Value
	limit := field.Arguments.ForName(LIMIT)
	if limit != nil {
		value = limit.Value
	} else {
		value = &ast.Value{
			Kind: ast.IntValue,
			Raw:  strconv.Itoa(my.meta.cfg.DefaultLimit),
		}
	}
	my.renderParam(value)
}

func (my *compilerContext) renderOffset(field *ast.Field) {
	offset := field.Arguments.ForName(OFFSET)
	if offset == nil {
		return
	}
	my.Write(` OFFSET `)
	my.renderParam(offset.Value)
}
