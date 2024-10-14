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
	symbols := []symbol{
		{"in", "Is in list of values"},
		{"eq", "Equals value"},
		{"is", "Is value null (true) or not null (false)"},
		{"ne", "Does not equal value"},
		{"gt", "Is greater than value"},
		{"lt", "Is lesser than value"},
		{"ge", "Is greater than or equals value"},
		{"le", "Is lesser than or equals value"},
		{"like", "Value matching (case-insensitive) pattern where '%' represents zero or more characters and '_' represents a single character. Eg. '_r%' finds values having 'r' in second position"},
		{"iLike", "Value matching (case-insensitive) pattern where '%' represents zero or more characters and '_' represents a single character. Eg. '_r%' finds values having 'r' in second position"},
		{"regex", "Value matches regex pattern"},
		{"iRegex", "Value matches (case-insensitive) regex pattern"},
	}
	operator := map[string][]symbol{
		SCALAR_ID:      symbols[:2],  //[in,eq]
		SCALAR_INT:     symbols[:8],  //[is,in,eq,ne,gt,ge,lt,le]
		SCALAR_FLOAT:   symbols[:8],  //[is,in,eq,ne,gt,ge,lt,le]
		SCALAR_STRING:  symbols,      //[is,in,eq,ne,gt,ge,lt,le,like,iLike,regex,iRegex]
		SCALAR_BOOLEAN: symbols[1:3], //[eq,is]
	}

	var build = func(scalar, suffix string, symbols []symbol) {
		name := util.JoinString(scalar, suffix)
		expr := &Class{Name: name, Kind: ast.InputObject, Fields: make(map[string]*Field)}
		for _, v := range symbols {
			t := ast.NamedType(scalar, nil)
			if v.Name == "is" {
				t = ast.NamedType(ENUM_IS_INPUT, nil)
			} else if v.Name == "in" {
				t = ast.ListType(ast.NonNullNamedType(scalar, nil), nil)
			}
			expr.Fields[v.Name] = &Field{Type: t, Name: v.Name, Description: v.Desc}
		}
		my.Nodes[name] = expr
	}
	for _, s := range scalars {
		build(s, SUFFIX_EXPRESSION, operator[s])
		build(s, SUFFIX_EXPRESSION_LIST, operator[s])
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
				"not": {
					Name: "not",
					Type: ast.NamedType(name, nil),
				},
				"and": {
					Name: "and",
					Type: ast.ListType(ast.NonNullNamedType(name, nil), nil),
				},
				"or": {
					Name: "or",
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
