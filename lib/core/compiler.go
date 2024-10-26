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
func (my *Compiler) Compile(operation *ast.OperationDefinition, variables json.RawMessage) (string, []any) {
	c := newContext(my.meta)
	c.Render(operation, variables)
	return c.String(), c.params
}
