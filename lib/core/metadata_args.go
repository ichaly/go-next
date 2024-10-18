package core

import (
	"github.com/duke-git/lancet/v2/slice"
	"github.com/ichaly/go-next/lib/util"
	"github.com/vektah/gqlparser/v2/ast"
)

type symbol struct {
	Name     string
	Text     string
	Describe string
}

func (my *Metadata) expression() error {
	var build = func(scalar, suffix string, symbols []*symbol) {
		name := util.JoinString(scalar, suffix)
		expr := &Class{Name: name, Kind: ast.InputObject, Fields: make(map[string]*Field)}
		for _, v := range symbols {
			var t *ast.Type
			if v.Name == IS {
				t = ast.NamedType(ENUM_IS_INPUT, nil)
			} else if v.Name == IN {
				t = ast.ListType(ast.NonNullNamedType(scalar, nil), nil)
			} else {
				t = ast.NamedType(scalar, nil)
			}
			expr.Fields[v.Name] = &Field{Type: t, Name: v.Name, Description: v.Describe}
		}
		my.Nodes[name] = expr
	}
	for _, s := range scalars {
		build(s, SUFFIX_EXPRESSION, symbols[s])
		build(s, SUFFIX_EXPRESSION_LIST, symbols[s])
	}
	return nil
}

func (my *Metadata) inputOption() error {
	for k, v := range my.Nodes {
		if v.Kind != ast.Object {
			continue
		}
		name := util.JoinString(k, SUFFIX_SORT_INPUT)
		sort := &Class{
			Name:   name,
			Kind:   ast.InputObject,
			Fields: make(map[string]*Field),
		}
		for _, f := range v.Fields {
			if !slice.Contain(scalars, f.Type.Name()) {
				continue
			}
			sort.Fields[f.Name] = &Field{
				Name: f.Name,
				Type: ast.NamedType(ENUM_SORT_INPUT, nil),
			}
		}
		my.Nodes[name] = sort
	}
	return nil
}

func (my *Metadata) whereOption() error {
	for k, v := range my.Nodes {
		if v.Kind != ast.Object {
			continue
		}
		name := util.JoinString(k, SUFFIX_WHERE_INPUT)
		where := &Class{
			Name: name,
			Kind: ast.InputObject,
			Fields: map[string]*Field{
				NOT: {
					Name: NOT,
					Type: ast.NamedType(name, nil),
				},
				AND: {
					Name: AND,
					Type: ast.ListType(ast.NonNullNamedType(name, nil), nil),
				},
				OR: {
					Name: OR,
					Type: ast.ListType(ast.NonNullNamedType(name, nil), nil),
				},
			},
		}
		for _, f := range v.Fields {
			if !slice.Contain(scalars, f.Type.Name()) {
				continue
			}
			where.Fields[f.Name] = &Field{
				Name: f.Name,
				Type: ast.NamedType(util.JoinString(f.Type.Name(), SUFFIX_EXPRESSION), nil),
			}
		}
		my.Nodes[name] = where
	}
	return nil
}
