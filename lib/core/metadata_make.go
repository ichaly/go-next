package core

import (
	"github.com/ichaly/go-next/lib/util"
	"github.com/vektah/gqlparser/v2/ast"
)

var inputs = func(name string) []*Input {
	return []*Input{
		{
			Name: DISTINCT,
			Type: ast.ListType(ast.NamedType(SCALAR_STRING, nil), nil),
		},
		{
			Name:    LIMIT,
			Type:    ast.NamedType(SCALAR_INT, nil),
			Default: `20`,
		},
		{
			Name: OFFSET,
			Type: ast.NamedType(SCALAR_INT, nil),
		},
		//{
		//	Name: FIRST,
		//	Type: ast.NamedType(SCALAR_INT, nil),
		//},
		//{
		//	Name: LAST,
		//	Type: ast.NamedType(SCALAR_INT, nil),
		//},
		//{
		//	Name: AFTER,
		//	Type: ast.NamedType(SCALAR_CURSOR, nil),
		//},
		//{
		//	Name: BEFORE,
		//	Type: ast.NamedType(SCALAR_CURSOR, nil),
		//},
		{
			Name: SORT,
			Type: ast.NamedType(util.JoinString(name, SUFFIX_SORT_INPUT), nil),
		},
		{
			Name: WHERE,
			Type: ast.NamedType(util.JoinString(name, SUFFIX_WHERE_INPUT), nil),
		},
	}
}

var mutationInputs = func(name string) []*Input {
	return []*Input{
		{
			Name: UPSERT,
			Type: ast.NamedType(util.JoinString(name, SUFFIX_UPSERT_INPUT), nil),
		},
		//{
		//	Name: INSERT,
		//	Type: ast.NamedType(util.JoinString(name, SUFFIX_INSERT_INPUT), nil),
		//},
		//{
		//	Name: UPDATE,
		//	Type: ast.NamedType(util.JoinString(name, SUFFIX_UPDATE_INPUT), nil),
		//},
		{
			Name: DELETE,
			Type: ast.NamedType(SCALAR_BOOLEAN, nil),
		},
	}
}

func (my *Metadata) entryOption() error {
	//构建入口节点
	query := &Class{Name: QUERY, Fields: make(map[string]*Field), Virtual: true}
	mutation := &Class{Name: MUTATION, Fields: make(map[string]*Field), Virtual: true}
	for k, v := range my.Nodes {
		if v.Kind != ast.Object {
			continue
		}
		_, name := my.Named(query.Name, k, JoinListSuffix())
		query.Fields[name] = &Field{
			Name:      name,
			Type:      ast.ListType(ast.NamedType(v.Name, nil), nil),
			Virtual:   query.Virtual,
			Arguments: inputs(k),
		}
		mutation.Fields[name] = &Field{
			Name:      name,
			Type:      ast.ListType(ast.NamedType(v.Name, nil), nil),
			Virtual:   mutation.Virtual,
			Arguments: append(inputs(k), mutationInputs(k)...),
		}
	}
	my.Nodes[query.Name] = query
	my.Nodes[mutation.Name] = mutation
	return nil
}
