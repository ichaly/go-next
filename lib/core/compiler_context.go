package core

import (
	"bytes"
	"github.com/duke-git/lancet/v2/convertor"
	stack "github.com/duke-git/lancet/v2/datastructure/stack"
	"github.com/vektah/gqlparser/v2/ast"
)

type compilerContext struct {
	buf  *bytes.Buffer
	meta *Metadata
}

type compilerElement struct {
	index int
	level int
	field *ast.Field
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
	sk := stack.NewLinkedStack[compilerElement]()
	my.WriteString(`SELECT jsonb_build_object(`)
	my.getFields(set, 0, func(index, level int, field *ast.Field) {
		size := sk.Size()
		if level == 0 {
			if index != 0 {
				my.WriteString(`,`)
			}
			my.WriteString(`'`)
			my.WriteString(field.Name)
			my.WriteString(`', __sj_`)
			my.WriteInt(index)
			my.WriteString(`.json`)
		}

		sk.Push(compilerElement{index: size, level: level, field: field})
	})
	my.WriteString(`) AS __root FROM (SELECT true) AS __root_x`)
	my.renderField(sk)
}

func (my *compilerContext) getFields(set ast.SelectionSet, level int, callback func(index, level int, field *ast.Field)) {
	for i, s := range set {
		switch t := s.(type) {
		case *ast.Field:
			_, ok := my.meta.Nodes[t.Definition.Type.Name()]
			if ok && callback != nil {
				callback(i, level, t)
				my.getFields(t.SelectionSet, level+1, callback)
			}
		}
	}
}

func (my *compilerContext) renderField(q *stack.LinkedStack[compilerElement]) {
	for {
		if q.IsEmpty() {
			break
		}
		e, _ := q.Pop()
		index, _, field := e.index, e.level, e.field

		my.renderJoin(index)
		my.renderList(index)
		my.renderJson(index)

		my.renderSelect(field)

		my.renderJsonClose(index)
		my.renderListClose(index)
		my.renderJoinClose(index)
	}
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

func (my *compilerContext) renderSelect(f *ast.Field) {
	node, ok := my.meta.Nodes[f.Definition.Type.Name()]
	if !ok {
		return
	}
	table, ok := my.meta.TableName(node.Name)
	if !ok {
		return
	}

	my.WriteString(`SELECT `)
	for i, s := range f.SelectionSet {
		switch typ := s.(type) {
		case *ast.Field:
			column, ok := my.meta.ColumnName(node.Name, typ.Name)
			if !ok {
				continue
			}
			if i != 0 {
				my.WriteString(",")
			}
			my.Quoted(table)
			my.WriteString(".")
			my.Quoted(column)
			my.WriteString(` AS `)
			my.Quoted(typ.Alias)
		}
	}
	my.WriteString(` FROM `)
	my.Quoted(table)
}
