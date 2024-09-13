package core

import (
	"github.com/duke-git/lancet/v2/maputil"
	"github.com/vektah/gqlparser/v2/ast"
	"strings"
)

var buildOption = func(my *Metadata) error {
	//构建节点信息
	for t, c := range my.tree {
		for k, f := range c.Fields {
			def := maputil.GetOrSet(my.Nodes, t, &ast.Definition{
				Name:        t,
				Description: f.TableDescription,
			})
			def.Fields = append(def.Fields, &ast.FieldDefinition{
				Name:        k,
				Type:        ast.NamedType(f.Type, nil),
				Description: f.Description,
			})
		}
	}
	//构建关联信息
	for _, v := range my.edge {
		for k, f := range v {
			currentTable, currentColumn := my.Named(
				f.Table, f.Name,
				WithTrimSuffix(),
				NamedRecursion(f, true),
			)
			foreignTable, foreignColumn := my.Named(
				f.TableRelation,
				f.ColumnRelation,
				WithTrimSuffix(),
				SwapPrimaryKey(currentTable),
				JoinListSuffix(),
				NamedRecursion(f, false),
			)
			println(k, currentTable, currentColumn, foreignTable, foreignColumn)
			//OneToMany
			my.Nodes[currentTable].Fields = append(my.Nodes[currentTable].Fields, &ast.FieldDefinition{
				Name: currentColumn,
				Type: ast.NamedType(foreignTable, nil),
			})
			//ManyToOne
			my.Nodes[foreignTable].Fields = append(my.Nodes[foreignTable].Fields, &ast.FieldDefinition{
				Name: foreignColumn,
				Type: ast.ListType(ast.NamedType(currentTable, nil), nil),
			})
			//如果是自关联的表则不进行多对多关联
			if f.Table == f.TableRelation {
				continue
			}
			//ManyToMany
			rest := maputil.OmitBy(v, func(key string, value Field) bool {
				return k == key || value.TableRelation == f.Table
			})
			for _, s := range rest {
				table, column := my.Named(
					s.TableRelation,
					s.Name,
					WithTrimSuffix(),
					JoinListSuffix(),
				)
				my.Nodes[foreignTable].Fields = append(my.Nodes[foreignTable].Fields, &ast.FieldDefinition{
					Name: column,
					Type: ast.ListType(ast.NamedType(table, nil), nil),
				})
			}
		}
	}
	return nil
}

var queryOption = func(my *Metadata) error {
	//构建Query
	query := &ast.Definition{Name: QUERY}
	for k, v := range my.Nodes {
		_, name := my.Named(query.Name, k, JoinListSuffix())
		query.Fields = append(query.Fields, &ast.FieldDefinition{
			Name: name,
			Arguments: []*ast.ArgumentDefinition{
				{
					Name: "id",
					Type: ast.NamedType("ID", nil),
				},
				{
					Name: "limit",
					Type: ast.NamedType("Int", nil),
				},
				{
					Name: "size",
					Type: ast.NamedType("Int", nil),
				},
				{
					Name: "first",
					Type: ast.NamedType("Int", nil),
				},
				{
					Name: "last",
					Type: ast.NamedType("Int", nil),
				},
				{
					Name: "after",
					Type: ast.NamedType("Cursor", nil),
				},
				{
					Name: "before",
					Type: ast.NamedType("Cursor", nil),
				},
				{
					Name: "distinctOn",
					Type: ast.ListType(ast.NamedType("String", nil), nil),
				},
				{
					Name: "sort",
					Type: ast.NamedType(strings.Join([]string{k, SUFFIX_SORT_INPUT}, ""), nil),
				},
				//{
				//	Name: "where",
				//	Type: ast.NamedType("WhereInput", nil),
				//},
			},
			Type: ast.ListType(ast.NamedType(v.Name, nil), nil),
		})
	}
	my.Nodes[query.Name] = query
	return nil
}
