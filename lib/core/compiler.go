package core

import (
	"encoding/json"
	"github.com/vektah/gqlparser/v2/ast"
)

type Compiler struct {
	meta   *Metadata
	schema *ast.Schema
}

func NewCompiler(m *Metadata, s *ast.Schema) *Compiler {
	return &Compiler{meta: m, schema: s}
}

func (my *Compiler) Compile(set ast.SelectionSet, vars json.RawMessage) (string, error) {
	c := newContext(my.meta)
	c.RenderQuery(set, vars)
	return c.String(), nil
}
