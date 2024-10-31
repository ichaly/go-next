package core

import (
	"github.com/vektah/gqlparser/v2/ast"
)

func (my *compilerContext) renderRemove(id, pid int, f *ast.Field) {
	remove := f.Arguments.ForName(REMOVE)
	if remove == nil || remove.Value.Raw != "true" {
		return
	}
	table, _ := my.meta.TableName(f.Definition.Type.Name(), false)
	my.Quoted(table)
	my.Space(`AS (DELETE FROM`)
	my.Quoted(table)
	my.renderWhereField(f)
	my.Space(`RETURNING`)
	my.Quoted(table)
	my.Write(`.* ) `)
}
