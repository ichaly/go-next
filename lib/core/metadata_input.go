package core

import (
	"github.com/vektah/gqlparser/v2/ast"
	"strings"
)

var inputOption = func(my *Metadata) error {
	for k, v := range my.Nodes {
		if v.Kind != ast.InputObject {
			continue
		}
		name := strings.Join([]string{k, SUFFIX_SORT_INPUT}, "")
		sort := &ast.Definition{
			Name: name,
			Kind: ast.InputObject,
		}
		for _, f := range v.Fields {
			sort.Fields = append(sort.Fields, &ast.FieldDefinition{
				Name: f.Name,
				Type: ast.NamedType("SortDirection", nil),
			})
		}
		my.Nodes[name] = sort
	}
	return nil
}
