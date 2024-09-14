package core

import (
	"github.com/ichaly/go-next/lib/util"
	"github.com/vektah/gqlparser/v2/ast"
)

func (my *Metadata) inputOption() error {
	for k, v := range my.Nodes {
		if v.Kind != ast.Object {
			continue
		}
		name := util.JoinString(k, SUFFIX_SORT_INPUT)
		sort := &ast.Definition{
			Name: name,
			Kind: ast.InputObject,
		}
		for _, f := range v.Fields {
			sort.Fields = append(sort.Fields, &ast.FieldDefinition{
				Name: f.Name,
				Type: ast.NamedType(ENUM_SORT_DIRECTION, nil),
			})
		}
		my.Nodes[name] = sort
	}
	return nil
}
