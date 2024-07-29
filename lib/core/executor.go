package core

import "github.com/vektah/gqlparser/v2/ast"

type Executor struct {
	s *ast.Schema
}

func NewExecutor(c *Compiler) *Executor { return &Executor{c} }
func (e *Executor) Execute() {
}
