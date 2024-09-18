package core

import (
	"github.com/duke-git/lancet/v2/slice"
	"github.com/ichaly/go-next/lib/util"
	"github.com/vektah/gqlparser/v2/ast"
)

var scalars = []string{SCALAR_ID, SCALAR_INT, SCALAR_FLOAT, SCALAR_STRING, SCALAR_BOOLEAN}

func (my *Metadata) expression() error {
	type symbol = struct {
		Name string
		Desc string
	}
	var build = func(scalar, suffix string, symbols []symbol) {
		name := util.JoinString(scalar, suffix)
		expr := &ast.Definition{Name: name, Kind: ast.InputObject}
		expr.Fields = append(expr.Fields, &ast.FieldDefinition{
			Name:        "isNull",
			Type:        ast.NamedType(SCALAR_BOOLEAN, nil),
			Description: "Is value null (true) or not null (false)",
		})
		for _, v := range symbols {
			expr.Fields = append(expr.Fields, &ast.FieldDefinition{
				Name: v.Name, Type: ast.NamedType(scalar, nil), Description: v.Desc,
			})
		}
		my.Nodes[name] = expr
	}
	for _, s := range scalars {
		build(s, SUFFIX_EXPR_LIST, []symbol{
			{"in", "Is in list of values"},
			{"notIn", "Is not in list of values"},
		})
		build(s, SUFFIX_EXPRESSION, []symbol{
			{"eq", "Equals value"},
			{"ne", "Does not equal value"},
			{"gt", "Is greater than value"},
			{"lt", "Is lesser than value"},
			{"ge", "Is greater than or equals value"},
			{"le", "Is lesser than or equals value"},
			{"like", "Value matching (case-insensitive) pattern where '%' represents zero or more characters and '_' represents a single character. Eg. '_r%' finds values having 'r' in second position"},
			{"notLike", "Value not matching pattern where '%' represents zero or more characters and '_' represents a single character. Eg. '_r%' finds values not having 'r' in second position"},
		})
	}
	return nil
}

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
			if !slice.Contain(scalars, f.Type.Name()) {
				continue
			}
			sort.Fields = append(sort.Fields, &ast.FieldDefinition{
				Name: f.Name,
				Type: ast.NamedType(ENUM_SORT_DIRECTION, nil),
			})
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
		where := &ast.Definition{
			Name: name,
			Kind: ast.InputObject,
			Fields: []*ast.FieldDefinition{
				{
					Name: "and",
					Type: ast.NamedType(name, nil),
				},
				{
					Name: "not",
					Type: ast.NamedType(name, nil),
				},
				{
					Name: "or",
					Type: ast.NamedType(name, nil),
				},
			},
		}
		for _, f := range v.Fields {
			if !slice.Contain(scalars, f.Type.Name()) {
				continue
			}
			where.Fields = append(where.Fields, &ast.FieldDefinition{
				Name: f.Name,
				Type: ast.NamedType(util.JoinString(f.Type.Name(), SUFFIX_EXPRESSION), nil),
			})
		}
		my.Nodes[name] = where
	}
	return nil
}
