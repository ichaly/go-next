package core

import (
	"github.com/ichaly/go-next/lib/core/internal/introspection"
	"github.com/ichaly/go-next/lib/util"
	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
	"gorm.io/gorm"
)

type Compiler struct {
	db     *gorm.DB
	meta   *Metadata
	schema *ast.Schema
}

func NewCompiler(m *Metadata, d *gorm.DB) (*Compiler, error) {
	input, err := m.MarshalSchema()
	if err != nil {
		return nil, err
	}
	s, err := gqlparser.LoadSchema(&ast.Source{Name: "schema", Input: input})
	if err != nil {
		return nil, err
	}
	return &Compiler{db: d, meta: m, schema: s}, nil
}

func (my *Compiler) Compile(query string) (interface{}, error) {
	doc, err := gqlparser.LoadQuery(my.schema, query)
	if err != nil {
		return nil, err
	}
	//doc.Operations.ForName()
	//IntrospectionQuery
	println(len(doc.Operations))
	return nil, nil
}

func (my *Compiler) Introspection() {
	i := introspection.NewIntrospection(my.schema)
	println(util.MustMarshalJson(i))
}
