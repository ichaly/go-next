package core

import (
	"github.com/vektah/gqlparser/v2/ast"
)

func (my *compilerContext) renderUpsert(id, pid int, f *ast.Field) {
	upsert := f.Arguments.ForName(UPSERT)
	where := f.Arguments.ForName(WHERE)
	if upsert == nil || where == nil {
		return
	}
	table, _ := my.meta.TableName(f.Definition.Type.Name(), false)

	my.Quoted(table)
	my.Space(`AS (INSERT INTO`)
	my.Quoted(table)
	my.Space(`(`)
	for i, v := range upsert.Value.Children {
		if i != 0 {
			my.Write(`,`)
		}
		my.Quoted(v.Name)
	}
	my.Write(`) SELECT `)
	for i, v := range upsert.Value.Children {
		if i != 0 {
			my.Write(`,`)
		}
		raw, _ := v.Value.Value(my.variables)
		my.Wrap(`'`, raw)
		my.Write(`::`)
		my.Write("text") //TODO:需要转化为数据库对应的具体类型
	}
	my.Space(`ON CONFLICT (id) DO UPDATE SET`)
	for i, v := range upsert.Value.Children {
		if i != 0 {
			my.Write(`,`)
		}
		my.Write(v.Name)
		my.Space(`=`)
		my.Write(`EXCLUDED.`)
		my.Write(v.Name)
	}
	my.renderWhereField(f)
	my.Space(`RETURNING`)
	my.Quoted(table)
	my.Write(`.* ) `)
}
