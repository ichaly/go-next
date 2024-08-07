package core

import (
	"context"
	"encoding/json"
	"github.com/dolmen-go/jsonmap"
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
		Variables     map[string]interface{} // raw variables from the JSON request
	}

	gqlResult struct {
		Data   jsonmap.Ordered `json:"data,omitempty"`
		Errors gqlerror.List   `json:"errors,omitempty"`
	}
)

type Executor struct {
	db     *gorm.DB
	meta   *Metadata
	intro  interface{}
	schema *ast.Schema
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
	return &Executor{db: d, meta: m, schema: s, intro: intro.New(s)}, nil
}

func (my *Executor) Execute(ctx context.Context, query string, vars json.RawMessage) (r gqlResult) {
	doc, err := gqlparser.LoadQuery(my.schema, query)
	if err != nil {
		r.Errors = err
		return
	}
	for _, d := range doc.Operations {
		println(d.Name)
		//op.GetSelections(ctx, d.SelectionSet, data, nil)
	}
	return
}
