package core

import (
	"bytes"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/ichaly/go-next/lib/util"
	"github.com/vektah/gqlparser/v2/ast"
)

type compilerContext struct {
	buf  *bytes.Buffer
	meta *Metadata
}

func newContext(m *Metadata) *compilerContext {
	return &compilerContext{meta: m, buf: bytes.NewBuffer([]byte{})}
}

func (my *compilerContext) String() string {
	return my.buf.String()
}

func (my *compilerContext) Quoted(identifier string) {
	my.buf.WriteByte('"')
	my.buf.WriteString(identifier)
	my.buf.WriteByte('"')
}

func (my *compilerContext) WriteInt(i int) {
	my.buf.WriteString(convertor.ToString(i))
}

func (my *compilerContext) WriteString(s string) {
	my.buf.WriteString(s)
}

func (my *compilerContext) RenderQuery(set ast.SelectionSet) {
	my.WriteString(`SELECT jsonb_build_object(`)
	my.eachField(set, func(index int, field *ast.Field) {
		if index != 0 {
			my.WriteString(`,`)
		}
		id := my.fieldFlag(field)
		my.WriteString(`'`)
		my.WriteString(field.Name)
		my.WriteString(`', __sj_`)
		my.WriteInt(id)
		my.WriteString(`.json`)
	})
	my.WriteString(`) AS __root FROM (SELECT true) AS __root_x`)
	my.renderField(0, set)
}

func (my *compilerContext) fieldFlag(field *ast.Field) int {
	p := field.GetPosition()
	return p.Line<<32 | p.Column
}

func (my *compilerContext) eachField(set ast.SelectionSet, callback func(index int, field *ast.Field)) {
	for i, s := range set {
		switch t := s.(type) {
		case *ast.Field:
			_, ok := my.meta.Nodes[t.Definition.Type.Name()]
			if ok && callback != nil {
				callback(i, t)
			}
		}
	}
}

func (my *compilerContext) renderField(pid int, set ast.SelectionSet) {
	my.eachField(set, func(index int, field *ast.Field) {
		id := my.fieldFlag(field)

		my.renderJoin(id)
		my.renderList(id)
		my.renderJson(id)

		my.renderSelect(id, pid, field)
		if len(field.SelectionSet) > 0 {
			my.renderField(id, field.SelectionSet)
		}

		my.renderJsonClose(id)
		my.renderListClose(id)
		my.renderJoinClose(id)
	})
}

func (my *compilerContext) renderJoin(id int) {
	my.WriteString(` LEFT OUTER JOIN LATERAL (`)
}

func (my *compilerContext) renderJoinClose(id int) {
	my.WriteString(`) AS __sj_`)
	my.WriteInt(id)
	my.WriteString(` ON true`)
}

func (my *compilerContext) renderList(id int) {
	my.WriteString(`SELECT COALESCE(jsonb_agg(__sj_`)
	my.WriteInt(id)
	my.WriteString(`.json), '[]') AS json FROM (`)
}

func (my *compilerContext) renderListClose(id int) {
	my.WriteString(`) AS __sj_`)
	my.WriteInt(id)
}

func (my *compilerContext) renderJson(id int) {
	my.WriteString(`SELECT to_jsonb(__sr_`)
	my.WriteInt(id)
	my.WriteString(`.*) AS json FROM ( `)
}

func (my *compilerContext) renderJsonClose(id int) {
	my.WriteString(`) AS __sr_`)
	my.WriteInt(id)
}

func (my *compilerContext) renderSelect(id, pid int, field *ast.Field) {
	name := field.Definition.Type.Name()
	table, ok := my.meta.TableName(name, false)
	if !ok {
		return
	}

	alias := util.JoinString(table, "_", convertor.ToString(id))

	my.WriteString(` SELECT `)
	for i, s := range field.SelectionSet {
		switch f := s.(type) {
		case *ast.Field:
			if i != 0 {
				my.WriteString(",")
			}
			if _, ok := my.meta.TableName(f.Definition.Type.Name(), false); ok {
				my.Quoted(util.JoinString("__sj_", convertor.ToString(my.fieldFlag(f))))
				my.WriteString(".")
				my.Quoted("json")
			} else {
				my.Quoted(alias)
				my.WriteString(".")
				my.Quoted(f.Name)
			}
			my.WriteString(` AS `)
			my.Quoted(f.Alias)
		}
	}
	my.WriteString(`FROM ( SELECT `)
	for i, s := range field.SelectionSet {
		switch f := s.(type) {
		case *ast.Field:
			column, ok := my.meta.ColumnName(name, f.Name, false)
			if !ok {
				continue
			}
			if i != 0 {
				my.WriteString(",")
			}
			my.Quoted(table)
			my.WriteString(".")
			my.Quoted(column)
		}
	}
	my.WriteString(` FROM `)
	my.Quoted(table)

	my.renderWhere(id, pid, field)

	my.WriteString(` LIMIT 20 ) AS`)
	my.Quoted(alias)
}

func (my *compilerContext) renderWhere(id, pid int, f *ast.Field) {
	class := my.meta.Nodes[f.ObjectDefinition.Name]
	field, ok := class.Fields[f.Name]
	if ok && len(field.Path) > 0 {
		my.WriteString(` WHERE (`)

		for i, v := range field.Path {
			if i != 0 {
				my.WriteString(" AND ")
			}

			my.Quoted(v.TableName)
			my.WriteString(".")
			my.Quoted(v.ColumnName)

			my.WriteString(" = ")

			my.Quoted(util.JoinString(v.TableRelation, "_", convertor.ToString(pid)))
			my.WriteString(".")
			my.Quoted(v.ColumnRelation)
		}

		my.WriteString(`)`)
	}
}
