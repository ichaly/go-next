package core

import (
	"fmt"
	"github.com/graphql-go/graphql"
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
	println(table.Name, field)

	for _, s := range p.Info.Operation.GetSelectionSet().Selections {
		fmt.Println(s)
	}
	p.Info.Schema.TypeMap()
	var w *strings.Builder
	_, _ = fmt.Fprintf(w, `SELECT json_object_agg('%s', %s) FROM (`, "", "")
	io.WriteString(w, `) AS "done_1337";`)

	return nil, nil
}
