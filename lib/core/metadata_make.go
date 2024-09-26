package core

import (
	"github.com/ichaly/go-next/lib/util"
	"github.com/vektah/gqlparser/v2/ast"
)

var inputs = func(name string) []*Input {
	return []*Input{
		{
			Name: "distinct",
			Type: ast.ListType(ast.NamedType(SCALAR_STRING, nil), nil),
		},
		{
			Name: "limit",
			Type: ast.NamedType(SCALAR_INT, nil),
		},
		{
			Name: "offset",
			Type: ast.NamedType(SCALAR_INT, nil),
		},
		{
			Name: "first",
			Type: ast.NamedType(SCALAR_INT, nil),
		},
		{
			Name: "last",
			Type: ast.NamedType(SCALAR_INT, nil),
		},
		{
			Name: "after",
			Type: ast.NamedType(SCALAR_CURSOR, nil),
		},
		{
			Name: "before",
			Type: ast.NamedType(SCALAR_CURSOR, nil),
		},
		{
			Name: "sort",
			Type: ast.NamedType(util.JoinString(name, SUFFIX_SORT_INPUT), nil),
		},
		{
			Name: "where",
			Type: ast.NamedType(util.JoinString(name, SUFFIX_WHERE_INPUT), nil),
		},
	}
}

func (my *Metadata) queryOption() error {
	//构建Query
	query := &Class{Name: QUERY, Fields: make(map[string]*Field), Virtual: true}
	for k, v := range my.Nodes {
		if v.Kind != ast.Object {
			continue
		}
		_, name := my.Named(query.Name, k, JoinListSuffix())
		query.Fields[name] = &Field{
			Name:    name,
			Type:    ast.ListType(ast.NamedType(v.Name, nil), nil),
			Virtual: query.Virtual,
			Arguments: append([]*Input{
				{
					Name: "id",
					Type: ast.NamedType(SCALAR_ID, nil),
				},
			}, inputs(k)...),
		}
	}
	my.Nodes[query.Name] = query
	return nil
}
