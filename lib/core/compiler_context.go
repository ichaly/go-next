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
			_, ok := my.meta.FindClass(t.Definition.Type.Name(), false)
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

func (my *compilerContext) renderSelect(id, pid int, f *ast.Field) {
	table, ok := my.meta.TableName(f.Definition.Type.Name(), false)
	if !ok {
		return
	}

	alias := util.JoinString(table, "_", convertor.ToString(id))

	my.WriteString(` SELECT `)
	for i, s := range f.SelectionSet {
		switch f := s.(type) {
		case *ast.Field:
			field, ok := my.meta.FindField(f.ObjectDefinition.Name, f.Name, false)
			if !ok {
				continue
			}
			if i != 0 {
				my.WriteString(",")
			}
			if len(field.Link) > 0 {
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
	my.WriteString(`FROM (`)
	field, ok := my.meta.FindField(f.Definition.Type.Name(), f.Name, false)
	if ok && field.Link == RECURSIVE {
		my.renderRecursiveSelect(id, pid, f)
	} else {
		my.renderUniversalSelect(id, pid, f)
	}

	my.WriteString(` LIMIT 20 ) AS`)
	my.Quoted(alias)
}

func (my *compilerContext) renderInner(id, pid int, f *ast.Field) {
	field, ok := my.meta.FindField(f.ObjectDefinition.Name, f.Name, false)
	if !ok || field.Join == nil {
		return
	}
	join := field.Join
	my.WriteString(` INNER JOIN `)
	my.WriteString(join.TableName)
	my.WriteString(` ON ((`)
	my.Quoted(join.TableName)
	my.WriteString(` . `)
	my.Quoted(join.ColumnName)
	my.WriteString(` = `)
	my.Quoted(util.JoinString(join.TableRelation, "_", convertor.ToString(pid)))
	my.WriteString(` . `)
	my.Quoted(join.ColumnRelation)
	my.WriteString(`))`)
}

func (my *compilerContext) renderWhere(id, pid int, f *ast.Field) {
	field, ok := my.meta.FindField(f.ObjectDefinition.Name, f.Name, false)
	if ok && field.Path != nil {
		path := field.Path
		my.WriteString(` WHERE (`)

		my.Quoted(path.TableName)
		my.WriteString(".")
		my.Quoted(path.ColumnName)

		my.WriteString(" = ")
		if field.Link == MANY_TO_MANY {
			my.Quoted(path.TableRelation)
		} else {
			my.Quoted(util.JoinString(path.TableRelation, "_", convertor.ToString(pid)))
		}
		my.WriteString(".")
		my.Quoted(path.ColumnRelation)

		my.WriteString(`)`)
	}
}

func (my *compilerContext) renderUniversalSelect(id, pid int, f *ast.Field) {
	table, _ := my.meta.TableName(f.Definition.Type.Name(), false)

	my.WriteString(` SELECT `)
	for i, s := range f.SelectionSet {
		switch f := s.(type) {
		case *ast.Field:
			_field, ok := my.meta.FindField(f.ObjectDefinition.Name, f.Name, false)
			if !ok || len(_field.Table) == 0 || len(_field.Column) == 0 {
				continue
			}
			if i != 0 {
				my.WriteString(",")
			}
			my.Quoted(_field.Table)
			my.WriteString(".")
			my.Quoted(_field.Column)
		}
	}
	my.WriteString(` FROM `)
	my.Quoted(table)

	my.renderInner(id, pid, f)
	my.renderWhere(id, pid, f)
}

func (my *compilerContext) renderRecursiveSelect(id, pid int, f *ast.Field) {
	field, _ := my.meta.FindField(f.Definition.Type.Name(), f.Name, false)
	table, column := field.Path.TableName, field.Path.ColumnName
	alias := util.JoinString("__rcte_", table)

	my.WriteString(` WITH RECURSIVE `)
	my.Quoted(alias)
	my.WriteString(` AS ((SELECT `)
	for i, s := range f.SelectionSet {
		switch f := s.(type) {
		case *ast.Field:
			_field, ok := my.meta.FindField(f.ObjectDefinition.Name, f.Name, false)
			if !ok || len(_field.Table) == 0 || len(_field.Column) == 0 {
				continue
			}
			if i != 0 {
				my.WriteString(",")
			}
			my.Quoted(_field.Table)
			my.WriteString(".")
			my.Quoted(_field.Column)
		}
	}

	my.WriteString(",")
	my.Quoted(table)
	my.WriteString(".")
	my.Quoted(column)
	my.WriteString(" FROM ")
	my.Quoted(table)

	my.WriteString(` LIMIT 1 ) UNION ALL `)

	my.WriteString(` SELECT `)
	for i, s := range f.SelectionSet {
		switch f := s.(type) {
		case *ast.Field:
			_field, ok := my.meta.FindField(f.ObjectDefinition.Name, f.Name, false)
			if !ok || len(_field.Table) == 0 || len(_field.Column) == 0 {
				continue
			}
			if i != 0 {
				my.WriteString(",")
			}
			my.Quoted(_field.Table)
			my.WriteString(".")
			my.Quoted(_field.Column)
			my.WriteString(" AS ")
			my.Quoted(_field.Column)
		}
	}

	my.WriteString(",")
	my.Quoted(table)
	my.WriteString(".")
	my.Quoted(column)
	my.WriteString(" FROM ")
	my.Quoted(table)

	my.WriteString(" , ")
	my.Quoted(alias)
	my.WriteString(") SELECT ")

	for i, s := range f.SelectionSet {
		switch f := s.(type) {
		case *ast.Field:
			_field, ok := my.meta.FindField(f.ObjectDefinition.Name, f.Name, false)
			if !ok || len(_field.Table) == 0 || len(_field.Column) == 0 {
				continue
			}
			if i != 0 {
				my.WriteString(",")
			}
			my.Quoted(_field.Table)
			my.WriteString(".")
			my.Quoted(_field.Column)
			my.WriteString(" AS ")
			my.Quoted(_field.Column)
		}
	}
	my.WriteString(` FROM (SELECT * FROM `)
	my.Quoted(alias)
	my.WriteString(` OFFSET 1) AS  `)
	my.Quoted(table)
}
