package core

import (
	"context"
	"encoding/json"
	"github.com/dolmen-go/jsonmap"
	"github.com/ichaly/go-next/lib/core/internal/intro"
	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type (
	gqlRequest struct {
		Query         string
		OperationName string
		Variables     map[string]interface{}
	}
	gqlResult struct {
		Sql    string          `json:"-"`
		Data   jsonmap.Ordered `json:"data,omitempty"`
		Errors gqlerror.List   `json:"errors,omitempty"`
	}
	gqlValue struct {
		value interface{}
		name  string
		err   error
	}
)

type Executor struct {
	intro    interface{}
	schema   *ast.Schema
	compiler *Compiler
}

func NewExecutor(c *Compiler, s *ast.Schema) (*Executor, error) {
	return &Executor{intro: intro.New(s), schema: s, compiler: c}, nil
}

func (my *Executor) Execute(ctx context.Context, query string, vars json.RawMessage) (r gqlResult) {
	doc, err := gqlparser.LoadQuery(my.schema, query)
	if err != nil {
		r.Errors = err
		return
	}
	//resultChans := make([]<-chan gqlValue, 0, len(set))
	for _, o := range doc.Operations {
		sql, _ := my.compiler.Compile(o.SelectionSet)
		r.Sql = sql
	}
	return
}
