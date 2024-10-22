package core

import (
	"encoding/json"
	"github.com/vektah/gqlparser/v2/ast"
)

type Compiler struct {
	meta *Metadata
}

func NewCompiler(m *Metadata) *Compiler {
	return &Compiler{meta: m}
}

func (my *Compiler) Compile(set ast.SelectionSet, vars json.RawMessage) (string, []any) {
	c := newContext(my.meta)
	c.RenderQuery(set, vars)
	return c.String(), c.params
}
