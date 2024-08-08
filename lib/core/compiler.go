package core

import (
	"bytes"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/vektah/gqlparser/v2/ast"
)

type Compiler struct {
	meta   *Metadata
	schema *ast.Schema
}

func NewCompiler(m *Metadata, s *ast.Schema) *Compiler {
	return &Compiler{meta: m, schema: s}
}

func (my *Compiler) Compile(set ast.SelectionSet) (string, error) {
	var w bytes.Buffer
	my.getSelection(set, &w)
	return w.String(), nil
}

func (my *Compiler) getSelection(set ast.SelectionSet, w *bytes.Buffer) {
	w.WriteString(`SELECT jsonb_build_object(`)

	for i, s := range set {
		switch typ := s.(type) {
		case *ast.Field:
			table, ok := my.meta.Nodes[typ.Definition.Type.Name()]
			if ok {
				if i != 0 {
					w.WriteString(`,`)
				}
				w.WriteString(`'`)
				w.WriteString(table.Original)
				w.WriteString(`', __sj_`)
				w.WriteString(convertor.ToString(i))
				w.WriteString(`.json`)
			}
			//println(typ.Definition.Type.Name(), typ.Name, table)
			//my.getSelection(typ.SelectionSet, w)
		}
	}

	w.WriteString(`) AS __root FROM ((SELECT true)) AS __root_x`)
}
