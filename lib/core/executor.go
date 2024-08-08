package core

import (
	"context"
	"encoding/json"
	"github.com/dolmen-go/jsonmap"
	"github.com/emirpasic/gods/stacks/arraystack"
	"github.com/ichaly/go-next/lib/core/internal/intro"
	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
)

type (
	gqlRequest struct {
		Query         string
		OperationName string
		Variables     map[string]interface{}
	}
	gqlResult struct {
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
	db       *gorm.DB
	meta     *Metadata
	intro    interface{}
	schema   *ast.Schema
	compiler *Compiler
}

func NewExecutor(m *Metadata, d *gorm.DB) (*Executor, error) {
	input, err := m.MarshalSchema()
	if err != nil {
		return nil, err
	}
	s, err := gqlparser.LoadSchema(&ast.Source{Name: "build", Input: input})
	if err != nil {
		return nil, err
	}
	return &Executor{db: d, meta: m, schema: s, intro: intro.New(s), compiler: NewCompiler(m, s)}, nil
}

func (my *Executor) Execute(ctx context.Context, query string, vars json.RawMessage) (r gqlResult) {
	doc, err := gqlparser.LoadQuery(my.schema, query)
	if err != nil {
		r.Errors = err
		return
	}
	stack := arraystack.New()
	//resultChans := make([]<-chan gqlValue, 0, len(set))
	for _, o := range doc.Operations {
		sql, _ := my.compiler.Compile(o.SelectionSet)
		println(sql)
	}
	println(stack.Size())
	return
}
