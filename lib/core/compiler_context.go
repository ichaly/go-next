package core

import (
	"bytes"
	"fmt"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/duke-git/lancet/v2/maputil"
	"github.com/ichaly/go-next/lib/util"
	"github.com/vektah/gqlparser/v2/ast"
	"strings"
)

type compilerContext struct {
	buf   *bytes.Buffer
	meta  *Metadata
	stack map[int]int
}

func newContext(m *Metadata) *compilerContext {
	return &compilerContext{meta: m, buf: bytes.NewBuffer([]byte{}), stack: make(map[int]int)}
}

func (my *compilerContext) String() string {
	return strings.TrimSpace(my.buf.String())
}

func (my *compilerContext) Quoted(elem ...any) *compilerContext {
	my.buf.WriteByte('"')
	my.WriteString(elem...)
	my.buf.WriteByte('"')
	return my
}

func (my *compilerContext) WriteString(elem ...any) *compilerContext {
	for _, e := range elem {
		my.buf.WriteString(fmt.Sprintf("%v", e))
	}
	return my
}

func (my *compilerContext) RenderQuery(set ast.SelectionSet) {
	my.WriteString(`SELECT jsonb_build_object(`)
	my.eachField(set, func(index int, field *ast.Field) {
		if index != 0 {
			my.WriteString(`,`)
		}
		id := my.fieldId(field)
		my.WriteString(`'`, field.Name, `', __sj_`, id, `.json`)
	})
	my.WriteString(`) AS __root FROM (SELECT true) AS __root_x`)
	my.renderField(0, set)
}

func (my *compilerContext) fieldId(field *ast.Field) int {
	p := field.GetPosition()
	id := p.Line<<32 | p.Column
	return maputil.GetOrSet(my.stack, id, len(my.stack))
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
		id := my.fieldId(field)

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
	my.WriteString(` ) AS `).Quoted(`__sj_`, id).WriteString(` ON true `)
}

func (my *compilerContext) renderList(id int) {
	my.WriteString(` SELECT COALESCE(jsonb_agg(__sj_`, id, `.json), '[]') AS json FROM ( `)
}

func (my *compilerContext) renderListClose(id int) {
	my.WriteString(` ) AS `).Quoted(`__sj_`, id)
}

func (my *compilerContext) renderJson(id int) {
	my.WriteString(` SELECT to_jsonb(__sr_`, id, `.*) AS json FROM ( `)
}

func (my *compilerContext) renderJsonClose(id int) {
	my.WriteString(` ) AS `).Quoted(`__sr_`, id)
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
			if len(field.Kind) > 0 {
				my.Quoted("__sj_", my.fieldId(f))
				my.WriteString(".").Quoted("json")
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
	if ok && field.Kind == RECURSIVE {
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
	if ok && field.Link != nil {
		path := field.Link
		my.WriteString(` WHERE (`)

		my.Quoted(path.TableName)
		my.WriteString(".")
		my.Quoted(path.ColumnName)

		my.WriteString(" = ")
		if field.Kind == MANY_TO_MANY {
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
	table, column := field.Link.TableName, field.Link.ColumnName
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

	my.WriteString(` WHERE `)
	my.Quoted(table).WriteString(".").WriteString(field.Link.ColumnRelation)
	my.WriteString(" = ")
	my.Quoted(table, "_", pid).WriteString(".").WriteString(field.Link.ColumnRelation)

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

	my.WriteString("WHERE (")
	my.WriteString("(").Quoted(table).WriteString(".").Quoted(column).WriteString("IS NOT NULL)")
	my.WriteString("AND").WriteString("(").Quoted(table).WriteString(".").Quoted(column).WriteString("!=").Quoted(field.Link.TableRelation).WriteString(".").Quoted(field.Link.ColumnRelation).WriteString(")")
	my.WriteString("AND").WriteString("(").Quoted(table).WriteString(".").Quoted(column).WriteString("=").Quoted(alias).WriteString(".").Quoted(field.Link.ColumnRelation).WriteString(")")
	my.WriteString(")")

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
