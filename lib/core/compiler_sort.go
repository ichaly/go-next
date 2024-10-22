package core

import (
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/vektah/gqlparser/v2/ast"
	"strings"
)

func (my *compilerContext) renderSort(field *ast.Field) {
	sort := field.Arguments.ForName(SORT)
	if sort == nil || len(sort.Value.Children) == 0 {
		return
	}
	my.Write(` ORDER BY  `)
	for i, v := range sort.Value.Children {
		if i != 0 {
			my.Write(`, `)
		}
		f, _ := my.meta.FindField(field.Definition.Type.Name(), v.Name, false)
		my.Quoted(f.Table)
		my.Write(".")
		my.Quoted(f.Column)
		my.Write(` `)
		my.Write(strings.ReplaceAll(convertor.ToString(v.Value.Raw), "_", " "))
	}
}
