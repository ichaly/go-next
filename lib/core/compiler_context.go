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
	my.Write(elem...)
	my.buf.WriteByte('"')
	return my
}

func (my *compilerContext) Write(elem ...any) *compilerContext {
	for _, e := range elem {
		my.buf.WriteString(fmt.Sprint(e))
	}
	return my
}

func (my *compilerContext) RenderQuery(set ast.SelectionSet) {
	my.Write(`SELECT jsonb_build_object(`)
	my.eachField(set, func(index int, field *ast.Field) {
		if index != 0 {
			my.Write(`,`)
		}
		id := my.fieldId(field)
		my.Write(`'`, field.Name, `', __sj_`, id, `.json`)
	})
	my.Write(`) AS __root FROM (SELECT true) AS __root_x`)
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
	my.Write(` LEFT OUTER JOIN LATERAL (`)
}

func (my *compilerContext) renderJoinClose(id int) {
	my.Write(` ) AS `).Quoted(`__sj_`, id).Write(` ON true `)
}

func (my *compilerContext) renderList(id int) {
	my.Write(` SELECT COALESCE(jsonb_agg(__sj_`, id, `.json), '[]') AS json FROM ( `)
}

func (my *compilerContext) renderListClose(id int) {
	my.Write(` ) AS `).Quoted(`__sj_`, id)
}

func (my *compilerContext) renderJson(id int) {
	my.Write(` SELECT to_jsonb(__sr_`, id, `.*) AS json FROM ( `)
}

func (my *compilerContext) renderJsonClose(id int) {
	my.Write(` ) AS `).Quoted(`__sr_`, id)
}

func (my *compilerContext) renderSelect(id, pid int, f *ast.Field) {
	table, ok := my.meta.TableName(f.Definition.Type.Name(), false)
	if !ok {
		return
	}

	alias := util.JoinString(table, "_", convertor.ToString(id))

	my.Write(` SELECT `)
	for i, s := range f.SelectionSet {
		switch f := s.(type) {
		case *ast.Field:
			field, ok := my.meta.FindField(f.ObjectDefinition.Name, f.Name, false)
			if !ok {
				continue
			}
			if i != 0 {
				my.Write(",")
			}
			if len(field.Kind) > 0 {
				my.Quoted("__sj_", my.fieldId(f))
				my.Write(".").Quoted("json")
			} else {
				my.Quoted(alias)
				my.Write(".")
				my.Quoted(f.Name)
			}
			my.Write(` AS `)
			my.Quoted(f.Alias)
		}
	}
	my.Write(`FROM (`)
	field, ok := my.meta.FindField(f.Definition.Type.Name(), f.Name, false)
	if ok && field.Kind == RECURSIVE {
		my.renderRecursiveSelect(id, pid, f)
	} else {
		my.renderUniversalSelect(id, pid, f)
	}

	my.Write(` LIMIT 20 ) AS`)
	my.Quoted(alias)
}

func (my *compilerContext) renderInner(id, pid int, f *ast.Field) {
	field, ok := my.meta.FindField(f.ObjectDefinition.Name, f.Name, false)
	if !ok || field.Join == nil {
		return
	}
	join := field.Join
	my.Write(` INNER JOIN `)
	my.Write(join.TableName)
	my.Write(` ON ((`)
	my.Quoted(join.TableName)
	my.Write(` . `)
	my.Quoted(join.ColumnName)
	my.Write(` = `)
	my.Quoted(util.JoinString(join.TableRelation, "_", convertor.ToString(pid)))
	my.Write(` . `)
	my.Quoted(join.ColumnRelation)
	my.Write(`))`)
}

func (my *compilerContext) renderWhere(id, pid int, f *ast.Field) {
	field, ok := my.meta.FindField(f.ObjectDefinition.Name, f.Name, false)
	where := f.Arguments.ForName("where")

	//TODO:处理关联关系的查询条件，有优化空间？
	if ok && field.Link != nil {
		link := field.Link
		if where == nil {
			where = &ast.Argument{Name: "where", Value: &ast.Value{Children: []*ast.ChildValue{}}}
		}
		var relation string
		if field.Kind == MANY_TO_MANY {
			relation = link.TableRelation
		} else {
			relation = util.JoinString(link.TableRelation, "_", convertor.ToString(pid))
		}
		where.Value.Children = append(where.Value.Children, &ast.ChildValue{
			Name: util.JoinString(`"`, link.TableName, `"."`, link.ColumnName, `"`),
			Value: &ast.Value{Children: []*ast.ChildValue{
				{Name: "eq", Value: &ast.Value{
					Children: []*ast.ChildValue{
						{Value: &ast.Value{
							Raw: util.JoinString(relation, ".", link.ColumnRelation),
						}},
					},
				}},
			}},
		})
	}

	//拼接SQL
	if where != nil {
		my.Write(` WHERE (`)
		my.renderWhereValue(where.Value)
		my.Write(`)`)
	}
}

func (my *compilerContext) renderUniversalSelect(id, pid int, f *ast.Field) {
	table, _ := my.meta.TableName(f.Definition.Type.Name(), false)

	my.Write(` SELECT `)
	for i, s := range f.SelectionSet {
		switch f := s.(type) {
		case *ast.Field:
			_field, ok := my.meta.FindField(f.ObjectDefinition.Name, f.Name, false)
			if !ok || len(_field.Table) == 0 || len(_field.Column) == 0 {
				continue
			}
			if i != 0 {
				my.Write(",")
			}
			my.Quoted(_field.Table)
			my.Write(".")
			my.Quoted(_field.Column)
		}
	}
	my.Write(` FROM `)
	my.Quoted(table)

	my.renderInner(id, pid, f)
	my.renderWhere(id, pid, f)
}

func (my *compilerContext) renderRecursiveSelect(id, pid int, f *ast.Field) {
	field, _ := my.meta.FindField(f.Definition.Type.Name(), f.Name, false)
	table, column := field.Link.TableName, field.Link.ColumnName
	alias := util.JoinString("__rcte_", table)

	my.Write(` WITH RECURSIVE `)
	my.Quoted(alias)
	my.Write(` AS ((SELECT `)
	for _, s := range f.SelectionSet {
		switch f := s.(type) {
		case *ast.Field:
			_field, ok := my.meta.FindField(f.ObjectDefinition.Name, f.Name, false)
			if !ok || len(_field.Table) == 0 || len(_field.Column) == 0 {
				continue
			}
			my.Quoted(_field.Table)
			my.Write(".")
			my.Quoted(_field.Column)
			my.Write(",")
		}
	}

	if "children" == f.Name {
		my.Quoted(field.Link.TableName).Write(".").Quoted(field.Link.ColumnName)
	} else {
		my.Quoted(field.Link.TableRelation).Write(".").Quoted(field.Link.ColumnRelation)
	}

	my.Write(" FROM ").Quoted(table).Write(` WHERE `)

	if "children" == f.Name {
		my.Quoted(table).Write(".").Write(field.Link.ColumnRelation)
		my.Write(" = ")
		my.Quoted(table, "_", pid).Write(".").Write(field.Link.ColumnRelation)
	} else {
		my.Quoted(table).Write(".").Write(field.Link.ColumnName)
		my.Write(" = ")
		my.Quoted(table, "_", pid).Write(".").Write(field.Link.ColumnName)
	}

	my.Write(` LIMIT 1 ) UNION ALL `)

	my.Write(` SELECT `)
	for i, s := range f.SelectionSet {
		switch f := s.(type) {
		case *ast.Field:
			_field, ok := my.meta.FindField(f.ObjectDefinition.Name, f.Name, false)
			if !ok || len(_field.Table) == 0 || len(_field.Column) == 0 {
				continue
			}
			if i != 0 {
				my.Write(",")
			}
			my.Quoted(_field.Table)
			my.Write(".")
			my.Quoted(_field.Column)
		}
	}

	if "children" == f.Name {
		my.Write(",").Quoted(table).Write(".").Quoted(field.Link.ColumnName).Write(" FROM ").Quoted(table)
	} else {
		my.Write(",").Quoted(table).Write(".").Quoted(field.Link.ColumnRelation).Write(" FROM ").Quoted(table)
	}

	my.Write(" , ")
	my.Quoted(alias)

	my.Write("WHERE (")
	if "children" == f.Name {
		my.Write("(").Quoted(table).Write(".").Quoted(column).Write("IS NOT NULL)")
		my.Write("AND").Write("(").Quoted(table).Write(".").Quoted(column).Write("!=").Quoted(field.Link.TableRelation).Write(".").Quoted(field.Link.ColumnRelation).Write(")")
		my.Write("AND").Write("(").Quoted(table).Write(".").Quoted(column).Write("=").Quoted(alias).Write(".").Quoted(field.Link.ColumnRelation).Write(")")
	} else {
		my.Write("(").Quoted(alias).Write(".").Quoted(field.Link.ColumnRelation).Write("IS NOT NULL)")
		my.Write("AND").Write("(").Quoted(alias).Write(".").Quoted(field.Link.ColumnRelation).Write("!=").Quoted(alias).Write(".").Quoted(field.Link.ColumnName).Write(")")
		my.Write("AND").Write("(").Quoted(alias).Write(".").Quoted(field.Link.ColumnRelation).Write("=").Quoted(field.Link.TableName).Write(".").Quoted(field.Link.ColumnName).Write(")")
	}
	my.Write(")")

	my.Write(") SELECT ")

	for i, s := range f.SelectionSet {
		switch f := s.(type) {
		case *ast.Field:
			_field, ok := my.meta.FindField(f.ObjectDefinition.Name, f.Name, false)
			if !ok || len(_field.Table) == 0 || len(_field.Column) == 0 {
				continue
			}
			if i != 0 {
				my.Write(",")
			}
			my.Quoted(_field.Table)
			my.Write(".")
			my.Quoted(_field.Column)
			my.Write(" AS ")
			my.Quoted(_field.Column)
		}
	}
	my.Write(` FROM (SELECT * FROM `)
	my.Quoted(alias)
	my.Write(` OFFSET 1) AS  `)
	my.Quoted(table)
}
