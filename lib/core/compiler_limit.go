package core

import (
	"github.com/vektah/gqlparser/v2/ast"
)

func (my *compilerContext) parseLimitValue(field *ast.Field) interface{} {
	limit := field.Arguments.ForName(LIMIT)
	if limit != nil {
		if value, err := limit.Value.Value(nil); err == nil {
			return value
		}
	}
	return my.meta.cfg.DefaultLimit
}

func (my *compilerContext) renderLimitField(field *ast.Field) {
	my.Write(` LIMIT `)
	my.Write(my.parseLimitValue(field))
}

func (my *compilerContext) renderOffsetField(field *ast.Field) {
	offset := field.Arguments.ForName(OFFSET)
	if offset == nil {
		return
	}
	value, err := offset.Value.Value(nil)
	if err != nil {
		return
	}
	my.Write(` OFFSET `)
	my.Write(value)
}
