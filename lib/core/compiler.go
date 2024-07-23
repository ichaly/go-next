package core

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"gorm.io/gorm"
	"io"
	"strings"
)

type Compiler struct {
	db   *gorm.DB
	meta *Metadata
}

func NewCompiler(m *Metadata, d *gorm.DB) (*Compiler, error) {
	return &Compiler{db: d, meta: m}, nil
}

func (my *Compiler) Compile(p graphql.ResolveParams) (interface{}, error) {
	field := p.Info.FieldName
	table := my.meta.Nodes[p.Info.ReturnType.Name()]
	input := string(p.Info.Operation.GetLoc().Source.Body)

	println(table.Name, field, input)

	my.recursion(p.Info.Operation.GetSelectionSet().Selections)

	var w *strings.Builder
	_, _ = fmt.Fprintf(w, `SELECT json_object_agg('%s', %s) FROM (`, "", "")
	io.WriteString(w, `) AS "done_1337";`)

	return nil, nil
}

func (my *Compiler) recursion(set []ast.Selection) {
	for _, s := range set {
		if s.GetSelectionSet() != nil && len(s.GetSelectionSet().Selections) > 0 {
			my.recursion(s.GetSelectionSet().Selections)
		}
		println("Field name is :", s.(*ast.Field).Name.Value)
	}
}
