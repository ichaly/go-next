package core

import "github.com/vektah/gqlparser/v2/ast"

func (my *compilerContext) renderInsert(id, pid int, f *ast.Field) {
	insert := f.Arguments.ForName(INSERT)
	table, _ := my.meta.TableName(f.Definition.Type.Name(), false)
	//field, _ := my.meta.FindField(f.Definition.Type.Name(), f.Name, false)

	my.Quoted(table)
	my.Write(` AS (`)
	my.Write(`INSERT INTO `)
	my.Quoted(table)

	my.Write(` (`)
	for i, v := range insert.Value.Children {
		if i != 0 {
			my.Write(`,`)
		}
		my.Quoted(v.Name)
	}

	my.Write(`) SELECT `)
	for i, v := range insert.Value.Children {
		if i != 0 {
			my.Write(`,`)
		}
		if val, err := v.Value.Value(my.variables); err == nil {
			my.Wrap(`'`, val)
			my.Write(`::`)
			my.Write("text")
		}
	}

	my.Write(` RETURNING `)
	my.Quoted(table)
	my.Write(`.* )`)
}
